package api

import (
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "message/api/qvbilam/message/v1"
)

type MessageServer struct {
	proto.UnimplementedMessageServer
}

func (s *MessageServer) CreatePrivateMessage(context.Context, *proto.CreatePrivateRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *MessageServer) CreateRoomMessage(context.Context, *proto.CreateRoomRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *MessageServer) CreateGroupMessage(context.Context, *proto.CreateGroupRequest) (*emptypb.Empty, error) {
	return nil, nil
}
