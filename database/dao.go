package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	dbHost     = "localhost"
	dbPort     = "5432"
	dbUser     = "postgres"
	dbName     = "postgres"
	dbPassword = "pass"
)

var _ UserDAO = (*DAO)(nil)

type DAO struct {
	db *gorm.DB
}

func CreateDAO(db *gorm.DB) *DAO {
	if err := db.AutoMigrate(&User{}).Error; err != nil {
		log.Printf("failed to migrate schema")
		return nil
	}

	return &DAO{
		db: db,
	}
}

func CreateConn() *gorm.DB {
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=false",
			dbHost, dbPort, dbUser, dbName, dbPassword))
	if err != nil {
		log.Println("open db error")
		return nil
	}

	return db
}

func (d *DAO) CreateUser(user CreateUserOption) error {
	err := d.db.Create(&User{Name: user.Name}).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *DAO) GetUser(user GetUserOption) (User, error) {
	panic("implement me")
}
