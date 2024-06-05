package application

import (
	"Job-service/internal/config"
	"Job-service/internal/database"
	"Job-service/internal/repository/mongo"
	"Job-service/internal/service"
	server "Job-service/internal/transport/grpc"
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

	jobRepository := mongo.NewJobRepository(mongoDatabase.Collection("job"))

	jobService := service.NewJobService(jobRepository)

	err = server.ServeGRPC(&cfg.GRPC, jobService)
	if err != nil {
		return fmt.Errorf("start gRPC server: %w", err)
	}

	log.Printf("gRPC started at %v\n", cfg.GRPC.Port)

	<-ctx.Done()

	return nil
}
