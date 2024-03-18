package main

import (
	"authservice/app/config"
	"authservice/app/database"
	"authservice/internal/delivery"
	"authservice/internal/repository"
	"authservice/internal/usecase"
	"log"
	"net"

	pb "authservice/internal/delivery/grpc"

	"google.golang.org/grpc"
	encrypts "authservice/helper/encrypt"
)

func main() {
	cfg := config.InitConfig()
	dbSql := database.InitDBMysql(cfg)
	hash := encrypts.New() 

	// Inisialisasi repository, usecase, dan delivery
	authRepo := repository.New(dbSql)
	authUsecase := usecase.New(authRepo, hash)
	authServiceServer := delivery.New(authUsecase)

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, authServiceServer) 

	// Start gRPC server
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
