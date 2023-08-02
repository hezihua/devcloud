package impl

import (
	"context"
	"hzh/devcloud/mcenter/apps/user"
)

// 查询用户列表
func (i *impl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (
	*user.User, error) {
	// 校验必填
	if err := in.Validate(); err != nil {
		return nil, err
	}
	ins := user.New(in)

	_, err := i.col.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 查询用户列表
func (i *impl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (
	*user.User, error) {
	return nil, nil
}

func (i *impl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (
	*user.User, error) {
	return nil, nil
}

// 查询用户列表
func (i *impl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (
	*user.UserSet, error) {
	return nil, nil
}

// 查询用户详情
func (i *impl) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (
	*user.User, error) {
	return nil, nil
}
