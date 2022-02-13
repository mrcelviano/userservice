package delivery

import (
	"context"
	"github.com/mrcelviano/userservice/internal/domain"
	server "github.com/mrcelviano/userservice/pkg/user/proto"
	"google.golang.org/grpc"
)

type userServer struct {
	service domain.UserService
}

func NewUserServer(service domain.UserService, opts ...grpc.ServerOption) *grpc.Server {
	u := &userServer{service: service}

	grpcServer := grpc.NewServer(opts...)
	server.RegisterUserServiceServer(grpcServer, u)
	return grpcServer
}

func (u *userServer) GetUserByID(ctx context.Context, req *server.GetUserByIDRequest) (resp *server.GetUserByIDResponse, err error) {
	user, err := u.service.GetByID(ctx, req.GetUserID())
	if err != nil {
		return &server.GetUserByIDResponse{}, err
	}
	return &server.GetUserByIDResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (u *userServer) SetIsRegisteredStatus(ctx context.Context, req *server.SetIsRegisteredStatusRequest) (resp *server.SetIsRegisteredStatusResponse, err error) {
	err = u.service.SetIsRegisteredStatus(ctx, req.GetUserID())
	if err != nil {
		return &server.SetIsRegisteredStatusResponse{}, err
	}
	return &server.SetIsRegisteredStatusResponse{StatusIsSet: true}, nil
}
