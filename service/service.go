package service

import (
	"context"
	"db-test/database"
	pb "db-test/pkg"
)

var _ (pb.DBServer) = (*Service)(nil)

type Service struct {
	dao *database.DAO
}

func CreateService(dao database.DAO) *Service {
	return &Service{
		dao: &dao,
	}
}

func (s *Service) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := s.dao.CreateUser(database.CreateUserOption{Name: in.Name})

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Service) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	resp, err := s.dao.GetUser(database.GetUserOption{Name: in.Name})

	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{Name: resp.Name}, nil
}
