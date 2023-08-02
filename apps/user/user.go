package user

import "hzh/devcloud/mcenter/common/meta"

func New(req *CreateUserRequest) *User {
	return &User{
		Meta: meta.NewMeta(),
		Spec: req,
	}
}
