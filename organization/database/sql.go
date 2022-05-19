package database

import (
	"fmt"
	"github.com/bimbimprasetyoafif/organization/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(user, pass, url, port, db string) (*gorm.DB, error) {
	fmt.Println("init database")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		url, user, pass, db, port)
	conn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	err = conn.AutoMigrate(&model.Organization{})
	if err != nil {
		return nil, err
	}
	fmt.Println("database ok")

	return conn, nil
}
