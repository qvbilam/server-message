package resource

import (
	"encoding/json"
	"message/enum"
)

type SystemObject struct {
	UserId      int64       `json:"user_id"`
	Object      string      `json:"object"`
	Type        string      `json:"type"`
	ContentType string      `json:"-"`
	Content     interface{} `json:"content"`
}

func (o *SystemObject) Encode() []byte {
	o.Type = enum.ObjTypeSystem
	body, _ := json.Marshal(o)
	return body
}

func (o *SystemObject) Decode(content []byte) SystemObject {
	obj := SystemObject{}
	err := json.Unmarshal(content, &obj)
	if err != nil {
		return SystemObject{}
	}
	return obj
}

type TipObject struct {
	UserId      int64       `json:"user_id"`
	Type        string      `json:"type"`
	ContentType string      `json:"-"`
	Content     interface{} `json:"content"`
}

func (o *TipObject) Encode() []byte {
	o.Type = enum.ObjTypeTips
	body, _ := json.Marshal(o)
	return body
}

func (o *TipObject) Decode(content []byte) TipObject {
	obj := TipObject{}
	err := json.Unmarshal(content, &obj)
	if err != nil {
		return TipObject{}
	}
	return obj
}

type PrivateObject struct {
	UserId      int64       `json:"user_id"`
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
	UserId      int64       `json:"user_id"`
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

type GroupObject struct {
	UserId      int64       `json:"user_id"`
	SendUserId  int64       `json:"send_user_id"`
	TargetId    int64       `json:"target_id"`
	Type        string      `json:"type"`
	ContentType string      `json:"-"`
	Content     interface{} `json:"content"`
}

func (o *GroupObject) Encode() []byte {
	o.Type = enum.ObjTypeGroup
	body, _ := json.Marshal(o)
	return body
}

func (o *GroupObject) Decode(content []byte) GroupObject {
	obj := GroupObject{}
	err := json.Unmarshal(content, &obj)
	if err != nil {
		return GroupObject{}
	}
	return obj
}
