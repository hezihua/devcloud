package user

import (
	context "context"

	grpc "google.golang.org/grpc"
)

const (
	AppName = "users"
)

type Service interface {
	// 创建用户
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// 更新用户
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	// 删除用户
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*User, error)

	RPCServer
}
