package service_test

import (
	"context"
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
	suite.svc = service.CreateService()
}

func (suite *ServiceTestSuite) TestCreate() {
	_, err := suite.svc.Create(context.Background(), &pb.CreateRequest{})
	assert.NoError(suite.T(), err)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}