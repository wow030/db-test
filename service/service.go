package service

import (
	"context"
	pb "db-test/pkg"
)

var _ (pb.DBServer) = (*Service)(nil)

type Service struct {
}

func CreateService() *Service {
	return &Service{}
}

func (s *Service) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateReply, error) {
	panic("implement me")
}
