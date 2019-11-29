package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var _ UserDAO = (*DAO)(nil)

type DAO struct {
	db *gorm.DB
}

func CreateDAO(db *gorm.DB) *DAO {

	return &DAO{
		db: db,
	}
}

func CreateConn() *gorm.DB {
	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword sslmode=false")
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
