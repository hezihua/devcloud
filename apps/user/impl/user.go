package impl

import (
	"context"
	"hzh/devcloud/mcenter/apps/user"
)

// 查询用户列表
func QueryUser(ctx context.Context, req *user.QueryUserRequest) (
	*user.UserSet, error) {
	return nil, nil
}

// 查询用户详情
func DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (
	*user.User, error) {
	return nil, nil
}
