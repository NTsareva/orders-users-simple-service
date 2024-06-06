package main

import (
	"context"
	"fmt"
	"net"

	"github.com/BurntSushi/toml"
	"github.com/NTsareva/orders-users-simple-service/internal/order-service/config"
	"github.com/NTsareva/orders-users-simple-service/internal/order-service/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NTsareva/orders-users-simple-service/order-service/ent"

	userproto "github.com/NTsareva/orders-users-simple-service/user-service/proto"

	orderproto "github.com/NTsareva/orders-users-simple-service/order-service/proto"
)

func createUserClient(config config.Config, logger *zap.Logger) userproto.UserServiceClient {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.UserService.Host, config.UserService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("did not connect to user-service", zap.Error(err))
	}
	return userproto.NewUserServiceClient(conn)
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var config config.Config
	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName)
	client, err := ent.Open("postgres", dataSourceName)
	if err != nil {
		logger.Fatal("failed opening connection to postgres", zap.Error(err))
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal("failed creating schema resources", zap.Error(err))
	}

	userClient := createUserClient(config, logger)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	s := grpc.NewServer()
	orderproto.RegisterOrderServiceServer(s, server.NewServer(client, logger, userClient))
	logger.Info("order-service started", zap.Int("port", config.Server.Port))
	if err := s.Serve(listener); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
