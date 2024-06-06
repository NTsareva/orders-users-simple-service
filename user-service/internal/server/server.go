package server

import (
	"go.uber.org/zap"

	"orders-users-simple-service/user-service/ent"
	"orders-users-simple-service/user-service/proto"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	Client *ent.Client
	Logger *zap.Logger
}
