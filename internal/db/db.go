package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	err    error
	dbHost = os.Getenv("DB_HOST")
	dbUser = os.Getenv("DB_USER")
	dbPort = os.Getenv("DB_PORT")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
)

func Init() error {
	dbTz := "America/Sao_Paulo"
	if os.Getenv("DB_TZ") != "" {
		dbTz = os.Getenv("DB_TZ")
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s TimeZone=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbTz, dbName)
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return err
	}

	log.Println("database connection established")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
