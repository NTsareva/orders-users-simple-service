package server

import (
	"go.uber.org/zap"

	"user-service/ent"
	"user-service/proto"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	Client *ent.Client
	Logger *zap.Logger
}
