package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	log.Info("テストログ")
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "テスト")
	})

	e.Logger.Fatal(e.Start(":8080"))
}