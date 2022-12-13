package business

import (
	"encoding/json"
	userProto "message/api/qvbilam/user/v1"
	"message/enum"
	"message/resource"
	"strconv"
)

type MessageBusiness struct {
	Type      string
	Content   string
	Url       string
	User      *userProto.UserResponse
	UserExtra string
	Extra     string
}

func (b *MessageBusiness) Resource() ([]byte, error) {

	switch b.Type {
	case enum.MsgTypeTxt: // 文本
		return b.textResource()
	case enum.MsgTypeImg: // 图片
		return b.imageResource()
	default: // 默认文本
		return b.textResource()
	}
}

func (b *MessageBusiness) textResource() ([]byte, error) {
	res := resource.Text{
		Content: b.Content,
		Extra:   b.Extra,
	}
	res.User = b.userResource()
	res.User.Extra = b.UserExtra
	return json.Marshal(res)
}

func (b *MessageBusiness) imageResource() ([]byte, error) {
	res := resource.Image{
		Content: b.Content,
		Url:     b.Url,
		Extra:   b.Extra,
	}

	res.User = b.userResource()
	res.User.Extra = b.UserExtra
	return json.Marshal(res)
}

func (b *MessageBusiness) userResource() resource.User {
	userId := int(b.User.Id)
	if b.User == nil {
		name := ""
		if userId == 0 {
			name = "system"
		}

		b.User = &userProto.UserResponse{
			Id:       b.User.Id,
			Nickname: name,
			Avatar:   "",
			Gender:   "",
		}
	}
	u := resource.User{
		Id:     strconv.Itoa(userId),
		Name:   b.User.Nickname,
		Avatar: b.User.Avatar,
		Extra:  b.UserExtra,
	}
	return u
}
