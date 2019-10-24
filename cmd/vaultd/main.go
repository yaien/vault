package main

import (
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/yaien/vault"
	"github.com/yaien/vault/pb"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterVaultServiceServer(server, &vault.VaultService{})
	reflection.Register(server)
	log.Println("grpc listening on :8080")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
