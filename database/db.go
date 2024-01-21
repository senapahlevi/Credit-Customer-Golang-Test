package database

import (
	"Credit-Customer-Golang-Test/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() (*Database, error) {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbnames := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASS")
	dbport := os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta", host, user, pass, dbnames, dbport)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, dbport, dbnames)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	errs := db.AutoMigrate(&models.Consumer{}, &models.Transaction{})
	if errs != nil {
		log.Fatal(errs)
	}

	return &Database{
		DB: db,
	}, nil
}
