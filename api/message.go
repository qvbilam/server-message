package api

import (
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "message/api/qvbilam/message/v1"
	"message/business"
	"message/enum"
	"message/global"
	"message/utils"
)

type MessageServer struct {
	proto.UnimplementedMessageServer
}

func (s *MessageServer) CreateQueue(ctx context.Context, request *proto.UpdateQueueRequest) (*proto.QueueResponse, error) {
	eName := request.ExchangeName
	if eName == "" {
		eName = global.ServerConfig.RabbitMQServerConfig.Exchange
	}

	qName := request.Name
	if qName == "" {
		qName = global.ServerConfig.RabbitMQServerConfig.QueuePrefix + utils.RandomCharacter(4)
	}
	b := business.QueueBusiness{ExchangeName: eName, Name: qName}
	entity, err := b.Create()
	if err != nil {
		return nil, err
	}
	return &proto.QueueResponse{
		Name:         entity.Name,
		ExchangeName: eName,
	}, nil
}

func (s *MessageServer) UpdateQueue(ctx context.Context, request *proto.UpdateQueueRequest) (*emptypb.Empty, error) {
	b := business.QueueBusiness{
		Name:   request.Name,
		Status: &request.Status,
	}
	if err := b.UpdateByName(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *MessageServer) DeleteQueue(ctx context.Context, request *proto.UpdateQueueRequest) (*emptypb.Empty, error) {
	closeStatus := enum.QueueStatusClose
	status := int64(closeStatus)
	b := business.QueueBusiness{
		Name:   request.Name,
		Status: &status,
	}
	if err := b.UpdateByName(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *MessageServer) CreatePrivateMessage(ctx context.Context, request *proto.CreatePrivateRequest) (*emptypb.Empty, error) {
	b := business.PrivateMessageBusiness{
		SenderUserId: request.UserId,
		TargetUserId: request.TargetUserId,
		ContentType:  request.Message.Type,
		Content: business.MessageBusiness{
			Code:    request.Message.Code,
			Type:    request.Message.Type,
			Content: request.Message.Content,
			Url:     request.Message.Url,
			Extra:   request.Message.Extra,
		},
	}
	_, err := b.CreateMessage()
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *MessageServer) CreateRoomMessage(ctx context.Context, request *proto.CreateRoomRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *MessageServer) CreateGroupMessage(ctx context.Context, request *proto.CreateGroupRequest) (*emptypb.Empty, error) {
	b := business.GroupMessageBusiness{
		SenderUserId:  request.UserId,
		TargetGroupId: request.GroupId,
		ContentType:   request.Message.Type,
		Content: business.MessageBusiness{
			Code:    request.Message.Code,
			Type:    request.Message.Type,
			Content: request.Message.Content,
			Url:     request.Message.Url,
			Extra:   request.Message.Extra,
		},
	}
	if _, err := b.CreateMessage(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *MessageServer) CreateGroupTxtMessage(ctx context.Context, request *proto.CreateGroupRequest) (*emptypb.Empty, error) {
	Type := enum.MsgTypeTxt
	b := business.GroupMessageBusiness{
		SenderUserId:  request.UserId,
		TargetGroupId: request.GroupId,
		ContentType:   Type,
		Content: business.MessageBusiness{
			Code:    request.Message.Code,
			Type:    Type,
			Content: request.Message.Content,
			Url:     request.Message.Url,
			Extra:   request.Message.Extra,
		},
	}
	if _, err := b.CreateMessage(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *MessageServer) CreateGroupCmdMessage(ctx context.Context, request *proto.CreateGroupRequest) (*emptypb.Empty, error) {
	Type := enum.CmdMsgType
	b := business.GroupMessageBusiness{
		SenderUserId:  request.UserId,
		TargetGroupId: request.GroupId,
		ContentType:   Type,
		Content: business.MessageBusiness{
			Code:    request.Message.Code,
			Type:    Type,
			Content: request.Message.Content,
			Url:     request.Message.Url,
			Extra:   request.Message.Extra,
		},
	}
	if _, err := b.CreateMessage(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *MessageServer) CreateGroupTipMessage(ctx context.Context, request *proto.CreateGroupRequest) (*emptypb.Empty, error) {
	Type := enum.TipMsgType
	b := business.GroupMessageBusiness{
		SenderUserId:  request.UserId,
		TargetGroupId: request.GroupId,
		ContentType:   Type,
		Content: business.MessageBusiness{
			Code:    request.Message.Code,
			Type:    Type,
			Content: request.Message.Content,
			Url:     request.Message.Url,
			Extra:   request.Message.Extra,
		},
	}
	if _, err := b.CreateMessage(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
