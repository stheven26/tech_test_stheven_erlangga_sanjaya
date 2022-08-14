package db

import (
	"fmt"

	"github.com/stheven26/config"
	"github.com/stheven26/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func SetupDB() {
	config := config.Configuration()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USERNAME, config.DB_PASSWORD, config.DB_NAME)

	// dsn := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@(" + config.DB_HOST + ")/" + config.DB_NAME + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Data{})
}

func GetConnection() *gorm.DB {
	return db
}
