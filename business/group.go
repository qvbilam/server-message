package business

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	contactProto "message/api/qvbilam/contact/v1"
	"message/global"
	"message/model"
	"message/resource"
)

type GroupMessageBusiness struct {
	SenderUserId  int64           `json:"sender_id"`
	TargetGroupId int64           `json:"target_id"`
	Type          string          `json:"type"`
	ContentType   string          `json:"-"`
	Content       MessageBusiness `json:"content"`
}

func (b *GroupMessageBusiness) CreateMessage() ([]byte, error) {
	sb := SenderBusiness{UserId: b.SenderUserId}
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
			UserID: b.SenderUserId,
		},
		GroupID:    b.TargetGroupId,
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
	b.send(mb)
	return m, nil
}

func (b *GroupMessageBusiness) send(mb MessageBusiness) {
	members, _ := global.ContactGroupServerClient.Members(context.Background(), &contactProto.SearchGroupMemberRequest{GroupId: b.TargetGroupId})
	for _, m := range members.Members {
		r := resource.GroupObject{
			UserId:      m.User.Id,
			SendUserId:  b.SenderUserId,
			TargetId:    b.TargetGroupId,
			ContentType: b.ContentType,
			Content:     mb,
		}

		body := r.Encode()
		fmt.Printf("content: %s\n", r.Content)
		fmt.Printf("body: %s\n", body)

		if err := PushDefaultExchange(body); err != nil {
			fmt.Printf("发送群聊聊消息失败:%s", err.Error())
		}
	}
}
