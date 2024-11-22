package main

import (
	"3-good-things/db"
	"3-good-things/models"
)

func main() {
	dbCon := db.NewDB()
	defer db.CloseDB(dbCon)
	dbCon.AutoMigrate(&models.GoodThings{})
}
