package main

import (
	"3-good-things/controller"
	"3-good-things/db"
	"3-good-things/repository"
	"3-good-things/usecase"
	"3-good-things/routes"
)

func main() {
	db := db.NewDB()

	GoodThingsRepository := repository.NewGoodThingsRepository(db)
	GoodThingsUsecase := usecase.NewGoodThingsUsecase(GoodThingsRepository)
	GoodThingsController := controller.NewGoodThingsController(GoodThingsUsecase)

	e := routes.NewGoodThingsRoutes(GoodThingsController)

	e.Logger.Fatal(e.Start(":8080"))
}
