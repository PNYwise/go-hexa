package handlers

import (
	"context"
	"go-hexa/internal/core/domain/enums/order"
	"go-hexa/internal/core/domain/models/requests"
	"go-hexa/internal/core/domain/services"
	"go-hexa/proto"
)

type userHandler struct {
	proto.UnimplementedUserServer
	userServie services.IUserService
}

func NewUserHandler(
	userServie services.IUserService,
) *userHandler {
	return &userHandler{
		userServie: userServie,
	}
}

func (u *userHandler) FindAll(ctx context.Context, request *proto.PaginationRequest) (*proto.UserResponses, error) {
	paginationRequest := &requests.PaginationRequest{
		Page:  int(request.GetPage()),
		Take:  int(request.GetTake()),
		Order: order.Enum(request.GetOrder().String()),
	}
	users, _ := u.userServie.FindAll(paginationRequest)
	protoResponses := make([]*proto.UserResponse, len(*users))

	for i, resp := range *users {
		protoResp := &proto.UserResponse{
			Id:      uint32(resp.ID),
			Name:    resp.Name,
			Email:   resp.Email,
			Address: resp.Address,
			Status:  string(resp.Status),
			Role:    string(resp.Role),
		}
		protoResponses[i] = protoResp
	}

	responses := &proto.UserResponses{
		UsersResponse: protoResponses,
	}
	return responses, nil
}

func (u *userHandler) FindOne(ctx context.Context, request *proto.FindOneRequest) (*proto.UserResponse, error) {
	user, err := u.userServie.FindOne(uint(request.GetUserId()))
	if err != nil {
		return nil, err
	}
	response := &proto.UserResponse{
		Id:      uint32(user.ID),
		Name:    user.Name,
		Email:   user.Name,
		Address: user.Address,
		Status:  string(user.Status),
		Role:    string(user.Role),
	}
	return response, nil
}
