package config

import (
	"errors"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")

	if dsn == "" {
		log.Fatal(errors.New("DB_URL environment variable is not set"))
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Gagal terhubung ke database", err)
	}
	log.Println("Berhasil terhubung ke database")

	return db
}
