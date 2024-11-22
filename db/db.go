package db

import (
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("DB接続成功")
	return db
}

func CloseDB(db *gorm.DB) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Error(err.Error())
	}
}
