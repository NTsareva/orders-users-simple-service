package server

import (
	"go.uber.org/zap"

	"github.com/NTsareva/orders-users-simple-service/cmd/user-service/ent"
	"github.com/NTsareva/orders-users-simple-service/cmd/user-service/proto"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	Client *ent.Client
	Logger *zap.Logger
}
