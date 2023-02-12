package business

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"message/global"
	"message/model"
	"message/resource"
)

type RoomMessageBusiness struct {
	SenderUserId int64           `json:"sender_id"`
	TargetRoomId int64           `json:"target_id"`
	Type         string          `json:"type"`
	ContentType  string          `json:"-"`
	Content      MessageBusiness `json:"content"`
}

func (b *RoomMessageBusiness) CreateMessage() ([]byte, error) {
	sb := SenderBusiness{
		LoginUserId: b.SenderUserId,
		UserId:      b.SenderUserId,
	}
	sender, err := sb.Sender()
	if err != nil {
		return nil, err
	}

	mb := MessageBusiness{
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

	entity := model.Room{
		UserModel: model.UserModel{
			UserID: b.SenderUserId,
		},
		RoomId:     b.TargetRoomId,
		MessageUid: messageEntity.Uid,
	}
	if res := global.DB.Save(&entity); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建私聊消息失败")
	}

	// 私聊类型
	r := resource.RoomObject{
		SendUserId:  b.SenderUserId,
		TargetId:    b.TargetRoomId,
		ContentType: b.ContentType,
		Content:     mb,
	}
	body := r.Encode()
	fmt.Printf("content: %s\n", r.Content)
	fmt.Printf("body: %s\n", body)

	if err := PushDefaultExchange(body); err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "发送私聊消息失败:%s", err.Error())
	}

	if err := PushChatRoomExchange(body); err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "发送私聊消息失败:%s", err.Error())
	}

	tx.Commit()
	return m, nil
}
