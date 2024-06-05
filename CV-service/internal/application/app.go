package application

import (
	"CV-service/internal/config"
	"CV-service/internal/database"
	"CV-service/internal/repository/mongo"
	"CV-service/internal/service"
	server "CV-service/internal/transport/grpc"
	"context"
	"fmt"
	"log"
)

func Run(ctx context.Context) error {
	cfg, err := config.Parse()
	if err != nil {
		return fmt.Errorf("build config: %w", err)
	}

	mongoDatabase, err := database.SetupMongoDatabase(ctx, &cfg.Mongo)
	if err != nil {
		return fmt.Errorf("start mongo connection: %w", err)
	}

	CVRepository := mongo.NewCVRepository(mongoDatabase.Collection("cv"))

	CVService := service.NewCVService(CVRepository)

	err = server.ServeGRPC(&cfg.GRPC, CVService)
	if err != nil {
		return fmt.Errorf("start gRPC server: %w", err)
	}

	log.Printf("gRPC started at %v\n", cfg.GRPC.Port)

	<-ctx.Done()
	log.Println("gRPC server shutting down")

	return nil
}
