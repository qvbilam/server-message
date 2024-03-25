package business

import (
	"context"
	contactProto "message/api/qvbilam/contact/v1"
	userProto "message/api/qvbilam/user/v1"
	"message/global"
)

type SenderBusiness struct {
	LoginUserId int64
	UserId      int64
	UserIds     []int64
}

func (b *SenderBusiness) Sender() (*SendUser, error) {
	var remark string

	// 获取好友备注
	users, _ := global.ContactFriendServerClient.Get(context.Background(), &contactProto.SearchFriendRequest{
		UserId:    b.LoginUserId,
		FriendIds: []int64{b.UserId},
	})
	if users != nil {
		if users.Total > 0 {
			for _, u := range users.Friends {
				remark = u.Remark
			}
		}
	}

	return b.protoToUser(b.UserId, nil, remark)
}

func (b *SenderBusiness) Senders() (map[int64]*SendUser, error) {
	senders, err := global.UserServerClient.List(context.Background(), &userProto.SearchRequest{Id: b.UserIds})
	if err != nil {
		return nil, err
	}
	friends, err := global.ContactFriendServerClient.Get(context.Background(), &contactProto.SearchFriendRequest{
		UserId:    b.LoginUserId,
		FriendIds: []int64{b.UserId},
	})

	res := make(map[int64]*SendUser)

	// 获取好友信息
	friendMap := make(map[int64]*contactProto.FriendResponse)
	if friends.Total > 0 {
		for _, u := range friends.Friends {
			friendMap[u.Friend.Id] = u
		}
	}

	// 获取用户信息
	userMap := make(map[int64]*userProto.UserResponse)
	if senders.Total > 0 {
		for _, u := range senders.Users {
			userMap[u.Id] = u
		}
	}

	for _, userId := range b.UserIds {
		var p *userProto.UserResponse
		var remark string
		if _, ok := userMap[userId]; ok {
			p = userMap[userId]
		}

		// 定义备注
		if _, ok := friendMap[userId]; ok {
			remark = friendMap[userId].Remark
		}
		res[userId], _ = b.protoToUser(userId, p, remark)
	}

	return res, nil
}

func (b *SenderBusiness) protoToUser(userId int64, user *userProto.UserResponse, remark string) (*SendUser, error) {
	if userId == 0 {
		return &SendUser{
			Id:       b.UserId,
			Code:     1000,
			Nickname: "系统",
			Avatar:   "",
		}, nil
	}
	var err error
	if user == nil {
		user, err = global.UserServerClient.Detail(context.Background(), &userProto.GetUserRequest{Id: b.UserId})
		if err != nil {
			return nil, err
		}
	}

	nickname := user.Nickname
	if remark != "" {
		nickname = remark
	}

	return &SendUser{
		Id:       user.Id,
		Code:     user.Code,
		Nickname: nickname,
		Avatar:   user.Avatar,
		Gender:   user.Gender,
		Extra:    nil,
	}, nil
}
