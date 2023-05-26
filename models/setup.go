package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//   dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	DBURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", DbHost, DbUser, DbPassword, DbName, DbPort)
	DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database")
	} else {
		fmt.Println("We are connected to the database")
	}
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&BookReq{})
	DB.AutoMigrate(&BookFood{})
	// DB.AutoMigrate(&OrderList{})
}
