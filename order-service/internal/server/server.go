package server

import (
	"context"
	"fmt"
	"strconv"

	"go.uber.org/zap"

	"orders-users-simple-service/order-service/ent"
	"orders-users-simple-service/order-service/ent/order"
	"orders-users-simple-service/order-service/proto"
	proto2 "user-service/proto"
)

type Server struct {
	proto.UnimplementedOrderServiceServer
	client     *ent.Client
	logger     *zap.Logger
	userClient proto2.UserServiceClient
}

// Function CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) is made for order creation
func (s *Server) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	order, err := s.client.Order.Create().SetID(strconv.Itoa(int(req.Order.Id))).
		SetTitle(req.Order.Title).
		SetDescription(req.Order.Description).
		SetUserID(int(req.Order.UserId)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create order: %v", err)
	}

	orderid, err := strconv.Atoi(order.ID)
	if err != nil {
		return nil, fmt.Errorf("parse order id: %v", err)
	}

	return &proto.OrderResponse{Order: &proto.Order{
		Id:          int32(orderid),
		Title:       order.Title,
		Description: order.Description,
		UserId:      int32(order.UserID),
	}, Message: "order created successfully"}, nil
}

// Function GetOrder(ctx context.Context, req *proto.GetOrderRequest) is made for get order
func (s *Server) GetOrder(ctx context.Context, req *proto.GetOrderRequest) (*proto.OrderResponse, error) {
	order, err := s.client.Order.
		Query().
		Where(order.ID(strconv.Itoa(int(req.Id)))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("error get order: %v", err)
	}

	orderid, err := strconv.Atoi(order.ID)
	if err != nil {
		return nil, fmt.Errorf("error parse order id: %v", err)
	}

	return &proto.OrderResponse{Order: &proto.Order{
		Id:          int32(orderid),
		Title:       order.Title,
		Description: order.Description,
		UserId:      int32(order.UserID),
	}, Message: "Order retrieved successfully"}, nil
}

// Function UpdateOrder(ctx context.Context, req *proto.UpdateOrderRequest) is made for update order
func (s *Server) UpdateOrder(ctx context.Context, req *proto.UpdateOrderRequest) (*proto.OrderResponse, error) {
	order, err := s.client.Order.
		UpdateOneID(strconv.Itoa(int(req.Order.Id))).
		SetTitle(req.Order.Title).
		SetDescription(req.Order.Description).
		SetUserID(int(req.Order.UserId)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("error update order: %v", err)
	}

	orderId, _ := strconv.Atoi(order.ID)
	if err != nil {
		return nil, fmt.Errorf("error parse order id: %v", err)
	}

	return &proto.OrderResponse{Order: &proto.Order{
		Id:          int32(orderId),
		Title:       order.Title,
		Description: order.Description,
		UserId:      int32(order.UserID),
	}, Message: "Order updated successfully"}, nil
}

// NewServer(client *ent.Client, logger *zap.Logger, userClient proto2.UserServiceClient) s implements new server
func NewServer(client *ent.Client, logger *zap.Logger, userClient proto2.UserServiceClient) *Server {
	return &Server{
		client:     client,
		logger:     logger,
		userClient: userClient,
	}
}
