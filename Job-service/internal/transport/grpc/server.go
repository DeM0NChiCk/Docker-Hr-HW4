package grpc

import (
	"Job-service/internal/config"
	"Job-service/internal/service"
	"Job-service/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func ServeGRPC(config *config.GRPCConfig, jobService *service.JobService) error {
	//addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	addr := fmt.Sprintf(":%s", config.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("start tcp listener: %w", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterJobServiceServer(grpcServer, jobService)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	return nil
}
