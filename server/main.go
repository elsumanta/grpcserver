package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/elsumanta/grpcserver/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//Import internal modules
	handler "github.com/elimSumanta/grpcserver/server/handler"
	repo "github.com/elimSumanta/grpcserver/server/repo"
)

var (
	port     = flag.Int("port", 50051, "The server port")
	dbhost   = "localhost"
	dbport   = 5656
	user     = "server"
	password = "server"
	dbname   = "grpc-server"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//init database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbhost, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	// Init repo modules
	repo.Init(db)

	// Init handler module
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterGreeterServer(s, handler.Init(context.Background()))

	// Serve port
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
