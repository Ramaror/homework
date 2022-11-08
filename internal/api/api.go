package api

import (
	"context"
	pb "homework/pkg/api"
)

func New() pb.AdminServer {
	return &implementation{}
}

type implementation struct {
	pb.UnimplementedAdminServer
}

func (i *implementation) UserCreate(ctx context.Context, in *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	return nil, nil
}
func (i *implementation) UserList(ctx context.Context, in *pb.UserListRequest) (*pb.UserListResponse, error) {
	return nil, nil
}
func (i *implementation) UserUpdate(ctx context.Context, in *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	return nil, nil
}
func (i *implementation) UserDelete(ctx context.Context, in *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	return nil, nil
}
