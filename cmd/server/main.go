package main

import (
	"context"
	"log"
	"net"

	vaultpb "go-secrets-vault/internal/vaultpb"

	"google.golang.org/grpc"
)

type vaultServer struct {
	vaultpb.UnimplementedVaultServer
	// todo: logic of vault
}

func (s *vaultServer) Store(ctx context.Context, req *vaultpb.Secret) (*vaultpb.SecretID, error) {
	// todo: cipher logic req.Ciphertext + req.Nonce -> save
	newID := "gend-uuid"
	log.Printf("Storing secret %s...", newID)
	return &vaultpb.SecretID{Id: newID}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("cannot start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	vaultpb.RegisterVaultServer(grpcServer, &vaultServer{})

	log.Println("server is up on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
