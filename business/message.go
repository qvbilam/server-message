package business

import (
	"encoding/json"
	"message/enum"
	"message/resource"
	"strconv"
)

type SendUser struct {
	Id       int64       `json:"id"`
	Code     int64       `json:"code"`
	Nickname string      `json:"nickname"`
	Avatar   string      `json:"avatar"`
	Gender   string      `json:"gender"`
	Extra    interface{} `json:"extra"`
}

type MessageBusiness struct {
	Code    int64       `json:"code"`
	Type    string      `json:"type"`
	Content string      `json:"content"`
	Url     string      `json:"url,omitempty"`
	User    *SendUser   `json:"user"`
	Extra   interface{} `json:"extra"`
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
		Code:    b.Code,
		Content: b.Content,
		Extra:   b.Extra,
	}
	res.Type = enum.MsgTypeTxt
	res.User = b.userResource()
	return json.Marshal(res)
}

func (b *MessageBusiness) imageResource() ([]byte, error) {
	res := resource.Image{
		Code:    b.Code,
		Content: b.Content,
		Url:     b.Url,
		Extra:   b.Extra,
	}
	res.Type = enum.MsgTypeImg
	res.User = b.userResource()

	return json.Marshal(res)
}

func (b *MessageBusiness) userResource() resource.User {
	userId := int(b.User.Id)
	if b.User == nil {
		name := ""
		if userId == 0 {
			name = "system"
		}

		b.User = &SendUser{
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
		Extra:  b.Extra,
	}
	return u
}
