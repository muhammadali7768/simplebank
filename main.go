package main

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muhammadali7768/simplebank/api"
	db "github.com/muhammadali7768/simplebank/db/sqlc"
	"github.com/muhammadali7768/simplebank/gapi"
	"github.com/muhammadali7768/simplebank/pb"
	"github.com/muhammadali7768/simplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	ctx := context.Background()

	connPool, err := pgxpool.New(ctx, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	store := db.NewStore(connPool)

	startGrpcServer(config, store)
	//startGinServer(config, store)

}

func startGrpcServer(config util.Config, store db.Store) {

	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create new server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("start gRPC at %s", lis.Addr().String())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Cannot start gRPC server:", err)
	}
}

func startGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create new server:", err)
	}

	err = server.Start(config.HttpServerAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
