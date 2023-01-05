package business

import (
	"context"
	userProto "message/api/qvbilam/user/v1"
	"message/global"
)

type SenderBusiness struct {
	UserId int64
}

func (b *SenderBusiness) Sender() (*SendUser, error) {
	if b.UserId == 0 {
		return &SendUser{
			Id:       b.UserId,
			Code:     1000,
			Nickname: "系统",
			Avatar:   "",
		}, nil
	}

	sender, err := global.UserServerClient.Detail(context.Background(), &userProto.GetUserRequest{Id: b.UserId})
	if err != nil {
		return nil, err
	}

	return &SendUser{
		Id:       sender.Id,
		Code:     sender.Code,
		Nickname: sender.Nickname,
		Avatar:   sender.Avatar,
		Gender:   sender.Gender,
		Extra:    nil,
	}, nil
}
