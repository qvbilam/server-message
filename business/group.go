package business

import (
	"context"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	contactProto "message/api/qvbilam/contact/v1"
	proto "message/api/qvbilam/message/v1"
	"message/enum"
	"message/global"
	"message/model"
	"message/resource"
)

type GroupMessageBusiness struct {
	UserId      int64           `json:"sender_id"`
	GroupId     int64           `json:"target_id"`
	Type        string          `json:"type"`
	ContentType string          `json:"-"`
	Content     MessageBusiness `json:"content"`

	Keyword string `json:"keyword"`

	Page    *int64 `json:"-"`
	PerPage *int64 `json:"-"`
}

func (b *GroupMessageBusiness) History() (int64, []*proto.MessageResponse, error) {
	var mRes []*proto.MessageResponse

	var userIds []int64
	total, ms := b.Messages()
	for _, m := range ms {
		userIds = append(userIds, m.UserID)
	}
	sb := GroupSenderBusiness{
		LoginUserId: b.UserId,
		GroupId:     b.GroupId,
		UserIds:     userIds,
	}
	senderMap, _ := sb.Senders()

	// content 转结构体
	for _, m := range ms {
		mb := MessageBusiness{}
		_ = json.Unmarshal([]byte(m.Message.Content), &mb)

		// 更新用户结构体
		if _, ok := senderMap[m.UserID]; ok {
			mb.User = senderMap[m.UserID]

			content, _ := mb.Resource()
			mRes = append(mRes, &proto.MessageResponse{
				UserId:      m.UserID,
				UID:         m.MessageUid,
				Type:        m.Type,
				Introduce:   m.Content,
				Content:     string(content),
				CreatedTime: m.CreatedAt.Unix(),
			})
		}
	}

	return total, mRes, nil
}

func (b *GroupMessageBusiness) Messages() (int64, []model.Group) {
	var count int64
	var m []model.Group
	condition := model.Group{
		UserModel: model.UserModel{
			UserID: b.UserId,
		},
		GroupID: b.GroupId,
	}

	// 只需要部分消息
	types := []string{enum.CmdMsgType, enum.TipMsgType}
	global.DB.Where(&condition).Where("type not in (?)", types).Model(&model.Group{}).Count(&count)
	if count == 0 {
		return 0, nil
	}

	page := *b.Page
	perPage := *b.PerPage

	if res := global.DB.Where("type not in (?)", types).
		Where(condition).
		Preload("Message").
		Order("id desc").
		Scopes(model.Paginate(int(page), int(perPage))).
		Find(&m); res.RowsAffected == 0 {
		return 0, nil
	}

	// 已读
	_, _ = global.ContactConversationServerClient.Read(context.Background(), &contactProto.UpdateConversationRequest{
		UserId:     b.UserId,
		ObjectType: enum.ObjTypeGroup,
		ObjectId:   b.GroupId,
	})

	return count, m
}

func (b *GroupMessageBusiness) CreateMessage() ([]byte, error) {
	sb := GroupSenderBusiness{
		LoginUserId: b.UserId,
		GroupId:     b.GroupId,
		UserId:      b.UserId,
	}
	sender, err := sb.Sender()
	if err != nil {
		return nil, err
	}

	mb := MessageBusiness{
		Code:    b.Content.Code,
		Type:    b.ContentType,
		Content: b.Content.Content,
		Url:     b.Content.Url,
		User:    sender,
		Extra:   b.Content.Extra,
	}
	m, err := mb.Resource()

	if err != nil {
		return nil, err
	}

	uId := uuid.NewV4()
	tx := global.DB.Begin()
	messageEntity := model.Message{
		Uid:     uId.String(),
		Type:    b.ContentType,
		Content: string(m),
	}
	if res := tx.Save(&messageEntity); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "保存私聊消息失败")
	}

	entity := model.Group{
		UserModel: model.UserModel{
			UserID: b.UserId,
		},
		GroupID:    b.GroupId,
		MessageUid: messageEntity.Uid,
		Type:       messageEntity.Type,
		Content:    b.Content.Content,
	}
	if res := global.DB.Save(&entity); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建私聊消息失败")
	}
	tx.Commit()
	// 发送群消息
	go func() {
		b.send(mb)
	}()
	return m, nil
}

func (b *GroupMessageBusiness) send(mb MessageBusiness) {
	members, _ := global.ContactGroupServerClient.Members(context.Background(), &contactProto.SearchGroupMemberRequest{GroupId: b.GroupId})
	if members == nil {
		return
	}

	r := resource.GroupObject{
		//UserId:      m.User.Id,
		SendUserId:  b.UserId,
		TargetId:    b.GroupId,
		ContentType: b.ContentType,
		Content:     mb,
	}

	for _, m := range members.Members {
		ur := r
		ur.UserId = m.User.Id

		body := ur.Encode()

		//fmt.Printf("content: %s\n", r.Content)
		//fmt.Printf("body: %s\n", body)

		if err := PushDefaultExchange(body); err != nil {
			fmt.Printf("队列发送群聊失败:%s", err.Error())
		}
	}

	if err := PushChatGroupExchange(r.Encode()); err != nil {
		fmt.Printf("队列发送群聊失败:%s", err.Error())
	}
}
