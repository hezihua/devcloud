package token

import context "context"

const (
	AppName = "tokens"
)

type Service interface {
	// 令牌颁发  restful
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	RPCServer
}

func NewValidateTokenRequirest(ak string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: ak,
	}
}
