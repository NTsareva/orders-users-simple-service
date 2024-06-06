package main

import (
	"fmt"
	"net"

	"github.com/BurntSushi/toml"
	config2 "github.com/NTsareva/orders-users-simple-service/internal/config-service/config"
	"github.com/NTsareva/orders-users-simple-service/internal/config-service/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"

	"github.com/NTsareva/orders-users-simple-service/user-service/ent"
	"github.com/NTsareva/orders-users-simple-service/user-service/proto"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	var config config2.Config
	if _, err := toml.DecodeFile("configs/config.toml", &config); err != nil {
		logger.Fatal("error loading config", zap.Error(err))
	}

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName)
	client, err := ent.Open("postgres", dataSourceName)
	if err != nil {
		logger.Fatal("error connecting to the database", zap.Error(err))
	}
	defer client.Close()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Database.Port))
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	serv := grpc.NewServer()
	proto.RegisterUserServiceServer(serv, &server.Server{Client: client, Logger: logger})
	logger.Info("user server started", zap.Int("port", config.Server.Port))
	if err := serv.Serve(listener); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
