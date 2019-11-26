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

func CreateDAO() DAO {
	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	if err != nil {
		log.Println("open db error")
		return DAO{}
	}

	return DAO{
		db: db,
	}
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
