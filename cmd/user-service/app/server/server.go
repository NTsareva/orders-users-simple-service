package server

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/NTsareva/orders-users-simple-service/cmd/user-service/ent"
	"github.com/NTsareva/orders-users-simple-service/cmd/user-service/ent/user"
	"github.com/NTsareva/orders-users-simple-service/cmd/user-service/proto"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	Client *ent.Client
	Logger *zap.Logger
}

func (s *Server) AddUser(ctx context.Context, req *proto.AddUserRequest) (*proto.UserResponse, error) {
	user := req.GetUser()

	newUser, err := s.Client.User.Create().
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetAge(int(user.Age)).
		Save(ctx)
	if err != nil {
		s.Logger.Error("failed to create user", zap.Error(err))
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &proto.UserResponse{
		User: &proto.User{
			Id:       int32(newUser.ID),
			Username: newUser.Username,
			Email:    newUser.Email,
			Age:      int32(newUser.Age),
		},
		Message: "user added successfully",
	}, nil
}

func (s *Server) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	email := req.GetEmail()

	user, err := s.Client.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		s.Logger.Error("failed to get user", zap.Error(err))
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &proto.UserResponse{
		User: &proto.User{
			Id:       int32(user.ID),
			Username: user.Username,
			Email:    user.Email,
			Age:      int32(user.Age),
		},
		Message: "user retrieved successfully",
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UserResponse, error) {
	user := req.GetUser()

	// Find user by ID
	existingUser, err := s.Client.User.Get(ctx, int(user.Id))
	if err != nil {
		s.Logger.Error("failed to find user", zap.Error(err))
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	updatedUser, err := existingUser.Update().
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetAge(int(user.Age)).
		Save(ctx)
	if err != nil {
		s.Logger.Error("failed to update user", zap.Error(err))
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return &proto.UserResponse{
		User: &proto.User{
			Id:       int32(updatedUser.ID),
			Username: updatedUser.Username,
			Email:    updatedUser.Email,
			Age:      int32(updatedUser.Age),
		},
		Message: "user updated successfully",
	}, nil
}
