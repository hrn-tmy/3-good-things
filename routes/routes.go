package routes

import (
	"3-good-things/controller"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

func NewGoodThingsRoutes(gtc controller.IGoodThingsController) *echo.Echo {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	log.Info("テストログ")
	e := echo.New()

	g := e.Group("/good-things")
	g.GET("", gtc.GetGoodThings)
	g.GET("/:id", gtc.GetGoodThingById)
	g.POST("", gtc.CreatedGoodThing)
	g.PUT("/:id", gtc.UpdateGoodThing)
	g.DELETE("/:id", gtc.DeleteGoodThing)

	return e
}
