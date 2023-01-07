// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.19.4
// source: conversation.proto

package contactV1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetConversationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetConversationRequest) Reset() {
	*x = GetConversationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConversationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConversationRequest) ProtoMessage() {}

func (x *GetConversationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConversationRequest.ProtoReflect.Descriptor instead.
func (*GetConversationRequest) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{0}
}

func (x *GetConversationRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Remark string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	IsDND  bool   `protobuf:"varint,5,opt,name=isDND,proto3" json:"isDND,omitempty"`
}

func (x *ObjectResponse) Reset() {
	*x = ObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectResponse) ProtoMessage() {}

func (x *ObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectResponse.ProtoReflect.Descriptor instead.
func (*ObjectResponse) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObjectResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ObjectResponse) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ObjectResponse) GetIsDND() bool {
	if x != nil {
		return x.IsDND
	}
	return false
}

type ConversationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          int64           `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ObjectType      string          `protobuf:"bytes,3,opt,name=objectType,proto3" json:"objectType,omitempty"`
	ObjectId        int64           `protobuf:"varint,4,opt,name=objectId,proto3" json:"objectId,omitempty"`
	Object          *ObjectResponse `protobuf:"bytes,5,opt,name=object,proto3" json:"object,omitempty"`
	NewsCount       int64           `protobuf:"varint,6,opt,name=newsCount,proto3" json:"newsCount,omitempty"`
	Tips            string          `protobuf:"bytes,7,opt,name=tips,proto3" json:"tips,omitempty"`
	LastMessage     string          `protobuf:"bytes,8,opt,name=lastMessage,proto3" json:"lastMessage,omitempty"`
	LastMessageTime int64           `protobuf:"varint,9,opt,name=lastMessageTime,proto3" json:"lastMessageTime,omitempty"`
}

func (x *ConversationResponse) Reset() {
	*x = ConversationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationResponse) ProtoMessage() {}

func (x *ConversationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationResponse.ProtoReflect.Descriptor instead.
func (*ConversationResponse) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{2}
}

func (x *ConversationResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConversationResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ConversationResponse) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *ConversationResponse) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *ConversationResponse) GetObject() *ObjectResponse {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *ConversationResponse) GetNewsCount() int64 {
	if x != nil {
		return x.NewsCount
	}
	return 0
}

func (x *ConversationResponse) GetTips() string {
	if x != nil {
		return x.Tips
	}
	return ""
}

func (x *ConversationResponse) GetLastMessage() string {
	if x != nil {
		return x.LastMessage
	}
	return ""
}

func (x *ConversationResponse) GetLastMessageTime() int64 {
	if x != nil {
		return x.LastMessageTime
	}
	return 0
}

type ConversationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total         int64                   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Conversations []*ConversationResponse `protobuf:"bytes,2,rep,name=conversations,proto3" json:"conversations,omitempty"`
	NewsCount     int64                   `protobuf:"varint,3,opt,name=newsCount,proto3" json:"newsCount,omitempty"`
}

func (x *ConversationsResponse) Reset() {
	*x = ConversationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationsResponse) ProtoMessage() {}

func (x *ConversationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationsResponse.ProtoReflect.Descriptor instead.
func (*ConversationsResponse) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{3}
}

func (x *ConversationsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ConversationsResponse) GetConversations() []*ConversationResponse {
	if x != nil {
		return x.Conversations
	}
	return nil
}

func (x *ConversationsResponse) GetNewsCount() int64 {
	if x != nil {
		return x.NewsCount
	}
	return 0
}

type UpdateConversationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ObjectType  string `protobuf:"bytes,3,opt,name=objectType,proto3" json:"objectType,omitempty"`
	ObjectId    int64  `protobuf:"varint,4,opt,name=objectId,proto3" json:"objectId,omitempty"`
	NewsCount   int64  `protobuf:"varint,5,opt,name=newsCount,proto3" json:"newsCount,omitempty"`
	Tips        string `protobuf:"bytes,6,opt,name=tips,proto3" json:"tips,omitempty"`
	LastMessage string `protobuf:"bytes,7,opt,name=lastMessage,proto3" json:"lastMessage,omitempty"`
	LastTime    int64  `protobuf:"varint,8,opt,name=lastTime,proto3" json:"lastTime,omitempty"`
}

func (x *UpdateConversationRequest) Reset() {
	*x = UpdateConversationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConversationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConversationRequest) ProtoMessage() {}

func (x *UpdateConversationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConversationRequest.ProtoReflect.Descriptor instead.
func (*UpdateConversationRequest) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateConversationRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateConversationRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateConversationRequest) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *UpdateConversationRequest) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *UpdateConversationRequest) GetNewsCount() int64 {
	if x != nil {
		return x.NewsCount
	}
	return 0
}

func (x *UpdateConversationRequest) GetTips() string {
	if x != nil {
		return x.Tips
	}
	return ""
}

func (x *UpdateConversationRequest) GetLastMessage() string {
	if x != nil {
		return x.LastMessage
	}
	return ""
}

func (x *UpdateConversationRequest) GetLastTime() int64 {
	if x != nil {
		return x.LastTime
	}
	return 0
}

var File_conversation_proto protoreflect.FileDescriptor

var file_conversation_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x30, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x7a, 0x0a, 0x0e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x44, 0x4e, 0x44,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44, 0x4e, 0x44, 0x22, 0xae, 0x02,
	0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x70, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x70, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x95,
	0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x48,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x65, 0x77,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xef, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x65, 0x77,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xf6, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x27, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x2a, 0x5a, 0x28, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x71, 0x76, 0x62, 0x69, 0x6c, 0x61, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conversation_proto_rawDescOnce sync.Once
	file_conversation_proto_rawDescData = file_conversation_proto_rawDesc
)

func file_conversation_proto_rawDescGZIP() []byte {
	file_conversation_proto_rawDescOnce.Do(func() {
		file_conversation_proto_rawDescData = protoimpl.X.CompressGZIP(file_conversation_proto_rawDescData)
	})
	return file_conversation_proto_rawDescData
}

var file_conversation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_conversation_proto_goTypes = []interface{}{
	(*GetConversationRequest)(nil),    // 0: contactPb.v1.GetConversationRequest
	(*ObjectResponse)(nil),            // 1: contactPb.v1.ObjectResponse
	(*ConversationResponse)(nil),      // 2: contactPb.v1.ConversationResponse
	(*ConversationsResponse)(nil),     // 3: contactPb.v1.ConversationsResponse
	(*UpdateConversationRequest)(nil), // 4: contactPb.v1.UpdateConversationRequest
	(*emptypb.Empty)(nil),             // 5: google.protobuf.Empty
}
var file_conversation_proto_depIdxs = []int32{
	1, // 0: contactPb.v1.ConversationResponse.object:type_name -> contactPb.v1.ObjectResponse
	2, // 1: contactPb.v1.ConversationsResponse.conversations:type_name -> contactPb.v1.ConversationResponse
	0, // 2: contactPb.v1.conversation.Get:input_type -> contactPb.v1.GetConversationRequest
	4, // 3: contactPb.v1.conversation.Create:input_type -> contactPb.v1.UpdateConversationRequest
	4, // 4: contactPb.v1.conversation.Delete:input_type -> contactPb.v1.UpdateConversationRequest
	3, // 5: contactPb.v1.conversation.Get:output_type -> contactPb.v1.ConversationsResponse
	5, // 6: contactPb.v1.conversation.Create:output_type -> google.protobuf.Empty
	5, // 7: contactPb.v1.conversation.Delete:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_conversation_proto_init() }
func file_conversation_proto_init() {
	if File_conversation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conversation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConversationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conversation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conversation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conversation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conversation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConversationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conversation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conversation_proto_goTypes,
		DependencyIndexes: file_conversation_proto_depIdxs,
		MessageInfos:      file_conversation_proto_msgTypes,
	}.Build()
	File_conversation_proto = out.File
	file_conversation_proto_rawDesc = nil
	file_conversation_proto_goTypes = nil
	file_conversation_proto_depIdxs = nil
}
