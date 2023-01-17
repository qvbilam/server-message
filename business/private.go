package business

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"message/enum"
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

	Page    *int64 `json:"-"`
	PerPage *int64 `json:"-"`
}

func (b *PrivateMessageBusiness) Messages() (int64, []model.Private) {
	var count int64
	var m []model.Private
	// 只需要部分消息
	types := []string{enum.CmdMsgType, enum.TipMsgType}

	global.DB.Model(&model.Private{}).Where("type not in (?)", types).Where(&model.Private{ChatSn: b.PrivateChatSn()}).Count(&count)
	if count == 0 {
		return 0, nil
	}

	if res := global.DB.Where("type not in (?)", types).Where(&model.Private{ChatSn: b.PrivateChatSn()}).Preload("Message").Find(&m); res.RowsAffected == 0 {
		return 0, nil
	}

	return count, m
}

func (b *PrivateMessageBusiness) CreateMessage() ([]byte, error) {
	sb := SenderBusiness{UserId: b.SenderUserId}
	sender, err := sb.Sender()
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

	fmt.Printf("private message business: %+v\n", b)

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
		ChatSn:       b.PrivateChatSn(),
		MessageUid:   messageEntity.Uid,
		Type:         messageEntity.Type,
		Content:      b.Content.Content,
	}
	if res := global.DB.Save(&entity); res.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建私聊消息失败")
	}

	// 私聊类型
	r := resource.PrivateObject{
		UserId:      b.TargetUserId,
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

func (b *PrivateMessageBusiness) PrivateChatSn() string {
	min := math.Min(float64(b.SenderUserId), float64(b.TargetUserId))
	max := math.Max(float64(b.SenderUserId), float64(b.TargetUserId))
	return fmt.Sprintf("%d-%d", int64(min), int64(max))
}
