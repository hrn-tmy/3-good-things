package main

import (
	"3-good-things/db"
	"3-good-things/models"
	"3-good-things/utils"
)

func main() {
	logger := utils.LoggerSetting()
	dbCon := db.NewDB()
	defer db.CloseDB(dbCon)
	dbCon.AutoMigrate(models.GoodThings{})
	logger.Info("マイグレート成功")
}
