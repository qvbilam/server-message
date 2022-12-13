package business

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	userProto "message/api/qvbilam/user/v1"
	"message/global"
	"message/model"
	"message/resource"
)

type PrivateMessageBusiness struct {
	SenderUserId int64           `json:"sender_id"`
	TargetUserId int64           `json:"target_id"`
	Type         string          `json:"type"`
	ContentType  string          `json:"-"`
	Content      MessageBusiness `json:"content"`
}

func (b *PrivateMessageBusiness) CreateMessage() ([]byte, error) {
	sender, err := global.UserServerClient.Detail(context.Background(), &userProto.GetUserRequest{Id: b.SenderUserId})
	if err != nil {
		return nil, err
	}

	mb := MessageBusiness{
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

	entity := model.Private{
		UserModel: model.UserModel{
			UserID: b.SenderUserId,
		},
		TargetUserId: b.TargetUserId,
		MessageUid:   messageEntity.Uid,
	}
	if res := global.DB.Save(&entity); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建私聊消息失败")
	}

	// 私聊类型
	r := resource.PrivateObject{
		SendUserId:  b.SenderUserId,
		TargetId:    b.TargetUserId,
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

	tx.Commit()
	return m, nil
}
