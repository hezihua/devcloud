package impl_test

import (
	"hzh/devcloud/mcenter/apps/token"
	"testing"
)

func TestIssueToken(t *testing.T) {
	// TODO 用户名密码错误，异常信息 自定义
	req := &token.IssueTokenRequest{
		Username: "admin",
		Password: "123456",
	}
	tk, err := impl.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequirest("cj5lkluigirrld45irj0")
	tk, err := impl.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}
