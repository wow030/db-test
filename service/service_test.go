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
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

const (
	dbHost     = "localhost"
	dbPort     = "5432"
	dbUser     = "postgres"
	dbName     = "postgres"
	dbPassword = "pass"
	maxRetry   = 10
)

type ServiceTestSuite struct {
	suite.Suite
	svc *service.Service
}

func (suite *ServiceTestSuite) SetupSuite() {

	var db *gorm.DB
	var err error
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)
	fmt.Println(connection)
	for i := 0; i < maxRetry; i++ {
		db, err = gorm.Open("postgres", connection)
		if err != nil {
			log.Println(err.Error())
			log.Printf("can't connect db, sleep %d second", i)
			time.Sleep(time.Duration(i) * time.Second)
		}
	}
	dao := database.CreateDAO(db)
	suite.svc = service.CreateService(dao)
}

func (suite *ServiceTestSuite) TestCreate() {
	_, err := suite.svc.CreateUser(context.Background(), &pb.CreateUserRequest{Name: "123123"})
	assert.NoError(suite.T(), err)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
