package service

import (
	"db-test/database"
	"context"
	pb "db-test/pkg"
)

var _ (pb.DBServer) = (*Service)(nil)

type Service struct {
	dao *database.DAO
}

func CreateService(dao database.DAO) *Service {
	return &Service{
		dao: &dao
	}
}

func (s *Service) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	panic("implement me")
}

func (s *Service) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserReply, error) {
	panic("implement me")
}
