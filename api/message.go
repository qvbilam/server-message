package api

import (
	"encoding/json"
	"go.uber.org/zap"
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

// CreateQueue 创建队列
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

// UpdateQueue 更新队列
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

// DeleteQueue 删除队列
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

// CreateSystemMessage 创建系统消息
func (s *MessageServer) CreateSystemMessage(ctx context.Context, request *proto.CreateSystemRequest) (*emptypb.Empty, error) {
	b := business.SystemMessageBusiness{
		Object:       request.Object,
		TargetUserId: request.UserId,
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

// CreateTipMessage 创建提示消息
func (s *MessageServer) CreateTipMessage(ctx context.Context, request *proto.CreateTipRequest) (*emptypb.Empty, error) {
	b := business.TipMessageBusiness{
		TargetUserId: request.UserId,
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

// CreatePrivateMessage 创建私聊消息
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

// CreateBroadcastUserMessage 创建广播用户消息
func (s *MessageServer) CreateBroadcastUserMessage(ctx context.Context, request *proto.CreateBroadcastUserRequest) (*emptypb.Empty, error) {
	userIds := request.UserIds
	if len(userIds) > 0 {
		for _, uId := range userIds {
			go func(userId int64) {
				b := business.TipMessageBusiness{
					TargetUserId: userId,
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
					c, _ := json.Marshal(&b)
					zap.S().Infof("广播用户消息: 失败: %s, content: %s", err, c)
				}
			}(uId)
		}
	}

	return &emptypb.Empty{}, nil
}

// CreateRoomMessage 创建房间消息
func (s *MessageServer) CreateRoomMessage(ctx context.Context, request *proto.CreateRoomRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// CreateGroupMessage 创建群组消息
func (s *MessageServer) CreateGroupMessage(ctx context.Context, request *proto.CreateGroupRequest) (*emptypb.Empty, error) {
	b := business.GroupMessageBusiness{
		UserId:      request.UserId,
		GroupId:     request.GroupId,
		ContentType: request.Message.Type,
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

// CreateGroupTxtMessage 创建群组文本消息
func (s *MessageServer) CreateGroupTxtMessage(ctx context.Context, request *proto.CreateGroupRequest) (*emptypb.Empty, error) {
	Type := enum.MsgTypeTxt
	b := business.GroupMessageBusiness{
		UserId:      request.UserId,
		GroupId:     request.GroupId,
		ContentType: Type,
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

// CreateGroupCmdMessage 创建文本命令消息
func (s *MessageServer) CreateGroupCmdMessage(ctx context.Context, request *proto.CreateGroupRequest) (*emptypb.Empty, error) {
	Type := enum.CmdMsgType
	b := business.GroupMessageBusiness{
		UserId:      request.UserId,
		GroupId:     request.GroupId,
		ContentType: Type,
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

// CreateGroupTipMessage 创建群组提示消息
func (s *MessageServer) CreateGroupTipMessage(ctx context.Context, request *proto.CreateGroupRequest) (*emptypb.Empty, error) {
	Type := enum.TipMsgType
	b := business.GroupMessageBusiness{
		UserId:      request.UserId,
		GroupId:     request.GroupId,
		ContentType: Type,
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

// GetPrivateMessage 获取私聊消息
func (s *MessageServer) GetPrivateMessage(ctx context.Context, request *proto.GetPrivateMessageRequest) (*proto.MessagesResponse, error) {
	b := business.PrivateMessageBusiness{
		SenderUserId: request.UserId,
		TargetUserId: request.TargetUserId,
		Keyword:      request.Keyword,
		Type:         request.Type,
		Page:         &request.Page.Page,
		PerPage:      &request.Page.PerPage,
	}

	count, messages, err := b.History()
	if err != nil {
		return nil, err
	}

	res := proto.MessagesResponse{Total: count, Messages: messages}
	return &res, nil
}

// GetGroupMessage 获取群组消息
func (s *MessageServer) GetGroupMessage(ctx context.Context, request *proto.GetGroupMessageRequest) (*proto.MessagesResponse, error) {
	b := business.GroupMessageBusiness{
		UserId:  request.UserId,
		GroupId: request.GroupId,
		Keyword: request.Keyword,
		Type:    request.Type,
		Page:    &request.Page.Page,
		PerPage: &request.Page.PerPage,
	}

	count, messages, err := b.History()
	if err != nil {
		return nil, err
	}
	res := proto.MessagesResponse{Total: count, Messages: messages}
	return &res, nil
}
