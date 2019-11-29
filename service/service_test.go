package service_test

import (
	"context"
	"db-test/database"
	pb "db-test/pkg"
	"db-test/service"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	svc *service.Service
}

func (suite *ServiceTestSuite) SetupSuite() {

	dao := database.CreateDAO(database.CreateConn())
	suite.svc = service.CreateService(dao)
}

func (suite *ServiceTestSuite) TestCreate() {
	_, err := suite.svc.CreateUser(context.Background(), &pb.CreateUserRequest{})
	assert.NoError(suite.T(), err)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
