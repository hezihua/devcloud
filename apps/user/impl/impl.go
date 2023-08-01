package impl

import (
	"hzh/devcloud/mcenter/apps/user"

	"github.com/infraboard/mcube/app"
	"google.golang.org/grpc"
)

var (
	svc = &impl{}
)

type impl struct {
	user.UnimplementedRPCServer
}

func (i *impl) Config() error {
	return nil
}

func (i *impl) Name() string {
	return user.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	user.RegisterRPCServer(server, i)
}

func init() {
	// grpc
	app.RegistryGrpcApp(svc)
	// internal app
	app.RegistryInternalApp(svc)

}
