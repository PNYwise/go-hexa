package main

import (
	"go-hexa/internal"
	"go-hexa/internal/core/configs"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	time.Local = time.UTC
	srv := grpc.NewServer()
	conf := configs.New()

	configs.ConnectDb()
	defer func() {
		if err := configs.CloseDb(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()
	internal.BootstrapGrpc(srv, configs.DB.Db, conf)

	port := conf.GetString("app.grpcport")
	log.Println("Starting RPC server at", ":"+port)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", ":"+port, err)
	}
	log.Println("server Started at", ":"+port)
	log.Fatal(srv.Serve(l))
}
