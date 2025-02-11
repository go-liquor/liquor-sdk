package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-liquor/liquor-sdk/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RegisterGRPCServer[T any, A any](implementation T, instance A, register func(imp T, registrar *grpc.Server)) fx.Option {
	return fx.Module("liquor-grpc-server", fx.Provide(instance), fx.Provide(func() *grpc.Server {
		return grpc.NewServer()
	}), fx.Invoke(register),
		fx.Invoke(
			func(cfg *config.Config, logger *zap.Logger, svc *grpc.Server, lc fx.Lifecycle) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						logger.Info("starting grpc server", zap.Int64("port", cfg.GetServerGrpcPort()))
						lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GetServerGrpcPort()))
						if err != nil {
							logger.Fatal("failed to start tcp grpc", zap.Error(err))
							return err
						}
						go svc.Serve(lis)
						return nil
					},
					OnStop: func(ctx context.Context) error {
						logger.Info("stopping grpc server")
						svc.GracefulStop()
						return nil
					},
				})

			},
		))
}
