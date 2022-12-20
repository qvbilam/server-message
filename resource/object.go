package resource

import (
	"encoding/json"
	"message/enum"
)

type PrivateObject struct {
	SendUserId  int64       `json:"send_user_id"`
	TargetId    int64       `json:"target_id"`
	Type        string      `json:"type"`
	ContentType string      `json:"-"`
	Content     interface{} `json:"content"`
}

func (o *PrivateObject) Encode() []byte {
	o.Type = enum.ObjTypePrivate
	body, _ := json.Marshal(o)
	return body
}

func (o *PrivateObject) Decode(content []byte) PrivateObject {
	obj := PrivateObject{}
	err := json.Unmarshal(content, &obj)
	if err != nil {
		return PrivateObject{}
	}
	return obj
}

type RoomObject struct {
	SendUserId  int64       `json:"send_user_id"`
	TargetId    int64       `json:"target_id"`
	Type        string      `json:"type"`
	ContentType string      `json:"-"`
	Content     interface{} `json:"content"`
}

func (o *RoomObject) Encode() []byte {
	o.Type = enum.ObjTypeRoom
	body, _ := json.Marshal(o)
	return body
}

func (o *RoomObject) Decode(content []byte) RoomObject {
	obj := RoomObject{}
	err := json.Unmarshal(content, &obj)
	if err != nil {
		return RoomObject{}
	}
	return obj
}