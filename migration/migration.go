package main

import (
	"3-good-things/db"
	"3-good-things/models"
	"log/slog"
	"os"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	dbCon := db.NewDB()
	defer db.CloseDB(dbCon)
	dbCon.AutoMigrate(models.GoodThings{})
	log.Info("マイグレート成功")
}
