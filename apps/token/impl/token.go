package impl

import (
	"context"
	"hzh/devcloud/mcenter/apps/token"
	"hzh/devcloud/mcenter/apps/token/provider"

	"go.mongodb.org/mongo-driver/bson"
)

func (i *impl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (
	*token.Token, error) {
	// 校验必填
	if err := in.Validate(); err != nil {
		return nil, err
	}

	issuer := provider.Get(in.GrantType)
	ins, err := issuer.IssueToken(ctx, in)
	if err != nil {
		return nil, err
	}

	_, err = i.col.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *impl) ValidateToken(ctx context.Context, in *token.ValidateTokenRequest) (
	*token.Token, error) {
	// 查找token
	tk := token.NewToken()
	err := i.col.FindOne(ctx, bson.M{
		"_id": in.AccessToken,
	}).Decode(tk)
	if err != nil {
		return nil, err
	}

	err = tk.CheckAliable()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
