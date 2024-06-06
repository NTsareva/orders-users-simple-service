package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/NTsareva/orders-users-simple-service/user-service/ent/user"
	"github.com/NTsareva/orders-users-simple-service/user-service/internal/server"
	"github.com/NTsareva/orders-users-simple-service/user-service/proto"

	"github.com/NTsareva/orders-users-simple-service/user-service/ent/enttest"
)

func TestAddUser(t *testing.T) {
	client := enttest.Open(t, "postgres", "host=localhost port=5432 user=postgres password=password dbname=userdb sslmode=disable")
	defer client.Close()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	srv := &server.Server{Client: client, Logger: logger}

	req := &proto.AddUserRequest{
		User: &proto.User{
			Id:       1,
			Username: "Test User",
			Email:    "test@example.com",
			Age:      25,
		},
	}

	res, err := srv.AddUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), res.User.Id)
	assert.Equal(t, "Test User", res.User.Username)
	assert.Equal(t, "test@example.com", res.User.Email)
	assert.Equal(t, int32(25), res.User.Age)

	usr, err := client.User.Query().Where(user.Email("test@example.com")).Only(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, usr.ID)
	assert.Equal(t, "Test User", usr.Username)
	assert.Equal(t, "test@example.com", usr.Email)
	assert.Equal(t, 25, usr.Age)
}
