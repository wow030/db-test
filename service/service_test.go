package service_test

import (
	"context"
	"db-test/database"
	pb "db-test/pkg"
	"db-test/service"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

const (
	dbHost     = "127.0.0.1"
	dbPort     = "5432"
	dbUser     = "postgres"
	dbName     = "postgres"
	dbPassword = "pass"
	maxRetry   = 5
)

type ServiceTestSuite struct {
	suite.Suite
	svc *service.Service
}

func (suite *ServiceTestSuite) SetupSuite() {

	var db *gorm.DB
	var err error
	for i := 1; i <= maxRetry; i++ {
		db, err = gorm.Open("postgres",
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=false",
				dbHost, dbPort, dbUser, dbName, dbPassword))
		if err != nil {
			log.Printf("can't create db, sleep %d second", i)
			time.Sleep(time.Duration(i) * time.Second)
		} else {
			break
		}

	}
	dao := database.CreateDAO(db)
	suite.svc = service.CreateService(dao)
}

func (suite *ServiceTestSuite) TestCreate() {
	_, err := suite.svc.CreateUser(context.Background(), &pb.CreateUserRequest{})
	assert.NoError(suite.T(), err)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
