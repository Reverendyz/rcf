package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/reverendyz/rcf/internal/utils"
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
	if len(dbHost) == 0 {
		return errors.New("DB_HOST environment variable must be set")
	}
	if len(dbUser) == 0 {
		return errors.New("DB_USER environment variable must be set")
	}
	if len(dbPort) == 0 {
		return errors.New("DB_PORT environment variable must be set")
	}
	if len(dbPass) == 0 {
		return errors.New("DB_PASS environment variable must be set")
	}
	if len(dbName) == 0 {
		return errors.New("DB_NAME environment variable must be set")
	}

	dbTz := utils.GetenvOrDefault("DB_TZ", "America/Sao_Paulo")
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
