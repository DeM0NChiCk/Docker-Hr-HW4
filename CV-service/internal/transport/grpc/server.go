package grpc

import (
	"CV-service/internal/config"
	"CV-service/internal/service"
	"CV-service/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func ServeGRPC(config *config.GRPCConfig, cvService *service.CVService) error {
	//addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	addr := fmt.Sprintf(":%s", config.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("start tcp listener: %w", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterCVServiceServer(grpcServer, cvService)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	return nil
}
