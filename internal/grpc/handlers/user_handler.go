package handlers

import (
	"context"
	"go-hexa/proto"
)

type UserServiceImpl struct {
	proto.UnimplementedUserServiceServer
}

func (u *UserServiceImpl) FindAll(context.Context, *proto.PaginationRequest) (*proto.UserResponses, error) {
	responses := &proto.UserResponses{
		UsersResponse: []*proto.UserResponse{
			{
				Id:      1,
				Name:    "bunayya",
				Email:   "bunayya@gmail.com",
				Address: "jln",
				Status:  proto.Status_ACTIVE,
				Role:    proto.Role_ADMIN,
			},
			{
				Id:      2,
				Name:    "wahyu",
				Email:   "wahyu@gmail.com",
				Address: "jln",
				Status:  proto.Status_ACTIVE,
				Role:    proto.Role_ADMIN,
			},
		},
	}
	return responses, nil
}

func (u *UserServiceImpl) FindOne(ctx context.Context, request *proto.FindOneRequest) (*proto.UserResponse, error) {
	response := new(proto.UserResponse)
	response.Id = request.GetUserId()
	return response, nil
}
