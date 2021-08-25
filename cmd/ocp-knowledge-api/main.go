package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	api "github.com/ozoncp/ocp-knowledge-api/internal/api"
	"github.com/ozoncp/ocp-knowledge-api/internal/repo"
	server "github.com/ozoncp/ocp-knowledge-api/pkg/ocp-knowledge-api"
	"google.golang.org/grpc"
)

const (
	grpcPort           = ":2021"
	grpcServerEndpoint = "localhost:2021"
)

func main() {
	fmt.Println("This is the 'ocp-knowledge-api' project'.")

	go runGrpcGateway()

	if err := runGrpcServer(); err != nil {
		log.Fatal(err)
	}
}

func runGrpcServer() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server.RegisterOcpKnowledgeApiServer(s, api.NewKnowledgeApi(initRepo()))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func initRepo() repo.KnowledgeRepo {
	dsn := "postgres://admin:admin@localhost:5432/knowledge_db?sslmode=disable"
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v\n", err)
	}

	log.Printf("Connection to database has been established.")

	return repo.NewRepo(db)
}

func runGrpcGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := server.RegisterOcpKnowledgeApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	log.Println("Server starting...")
	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}
