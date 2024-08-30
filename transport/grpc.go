package transport

import (
	"context"
	"net"

	"github.com/syoi-org/judy/internal/broker/pb"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcConfig struct {
	ListenAddr string `mapstructure:"listen_addr" yaml:"listen_addr" validate:"required"`
}

type GrpcParams struct {
	fx.In
	Config *GrpcConfig
	Broker pb.BrokerServer
}

func NewGrpc(p GrpcParams) *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterBrokerServer(grpcServer, p.Broker)
	reflection.Register(grpcServer)
	return grpcServer
}

func runGrpcServer(lifecycle fx.Lifecycle, grpcServer *grpc.Server, grpcConfig *GrpcConfig) error {
	listener, err := net.Listen("tcp", grpcConfig.ListenAddr)
	if err != nil {
		return err
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go grpcServer.Serve(listener)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})
	return nil
}
