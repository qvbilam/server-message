package business

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	contactProto "message/api/qvbilam/contact/v1"
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

	if res := global.DB.Where("type not in (?)", types).Where(condition).Preload("Message").Find(&m); res.RowsAffected == 0 {
		return 0, nil
	}

	return count, m
}

func (b *GroupMessageBusiness) CreateMessage() ([]byte, error) {
	sb := SenderBusiness{UserId: b.UserId}
	sender, err := sb.Sender()
	if err != nil {
		return nil, err
	}

	mb := MessageBusiness{
		Code:    b.Content.Code,
		Type:    b.ContentType,
		Content: b.Content.Content,
		Url:     b.Content.Url,
		User: &SendUser{
			Id:       sender.Id,
			Code:     sender.Code,
			Nickname: sender.Nickname,
			Avatar:   sender.Avatar,
			Gender:   sender.Gender,
			Extra:    "",
		},
		Extra: b.Content.Extra,
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

	for _, m := range members.Members {
		r := resource.GroupObject{
			UserId:      m.User.Id,
			SendUserId:  b.UserId,
			TargetId:    b.GroupId,
			ContentType: b.ContentType,
			Content:     mb,
		}

		body := r.Encode()
		//fmt.Printf("content: %s\n", r.Content)
		//fmt.Printf("body: %s\n", body)

		go func() {
			if mb.Type == enum.MsgTypeTxt {
				_, _ = global.ContactConversationServerClient.Create(context.Background(), &contactProto.UpdateConversationRequest{
					UserId:      m.User.Id,
					ObjectType:  enum.ObjTypeGroup,
					ObjectId:    b.GroupId,
					LastMessage: mb.Content,
				})
			}
		}()

		if err := PushDefaultExchange(body); err != nil {
			fmt.Printf("队列发送群聊失败:%s", err.Error())
		}

		if err := PushChatGroupExchange(body); err != nil {
			fmt.Printf("队列发送群聊消息失败:%s", err.Error())
		}
	}
}
