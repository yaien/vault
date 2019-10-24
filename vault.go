package vault

import (
	"context"

	"github.com/yaien/vault/pb"
	"golang.org/x/crypto/bcrypt"
)

type VaultService struct {
}

func (s *VaultService) Hash(ctx context.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	password := []byte(r.GetPassword())
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &pb.HashResponse{Hash: string(hash)}, nil
}

func (s *VaultService) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	password := []byte(r.GetPassword())
	hash := []byte(r.GetHash())
	err := bcrypt.CompareHashAndPassword(hash, password)
	return &pb.ValidateResponse{Valid: err == nil}, nil
}
