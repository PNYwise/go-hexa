package main

import (
	"go-hexa/internal/core/configs"
	"go-hexa/internal/grpc/handlers"
	"go-hexa/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	srv := grpc.NewServer()
	conf := configs.New()

	userSrv := &handlers.UserServiceImpl{}
	proto.RegisterUserServiceServer(srv, userSrv)

	port := conf.GetString("app.grpcport")
	log.Println("Starting RPC server at", ":"+port)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", ":"+port, err)
	}
	log.Println("server Started at", ":"+port)
	log.Fatal(srv.Serve(l))
}
