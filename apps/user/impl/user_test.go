package impl_test

import (
	"hzh/devcloud/mcenter/apps/user"
	"hzh/devcloud/mcenter/test/tools"
	"testing"
)

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = 	"123456"
	ins, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestQueryUser(t *testing.T) {
	impl.CreateUser(ctx, nil)
}

func TestDescribeUser(t *testing.T) {
	impl.CreateUser(ctx, nil)
}

func TestUpdateUser(t *testing.T) {
	impl.CreateUser(ctx, nil)
}

func TestDeleteUser(t *testing.T) {
	impl.CreateUser(ctx, nil)
}
