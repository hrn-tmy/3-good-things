package controller

import (
	"3-good-things/models"
	"3-good-things/usecase"
	"3-good-things/validation"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type IGoodThingsController interface {
	GetGoodThings(ctx echo.Context) error
	GetGoodThingById(ctx echo.Context) error
	CreatedGoodThing(ctx echo.Context) error
	UpdateGoodThing(ctx echo.Context) error
	DeleteGoodThing(ctx echo.Context) error
}

type GoodThingsController struct {
	gtu usecase.IGoodThingsUsecase
}

func NewGoodThingsController(gtu usecase.IGoodThingsUsecase) IGoodThingsController {
	return &GoodThingsController{gtu}
}

// 一覧取得コントローラー
func (gtc *GoodThingsController) GetGoodThings(ctx echo.Context) error {
	res, err := gtc.gtu.GetGoodThings()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

// 詳細取得コントローラー
func (gtc *GoodThingsController) GetGoodThingById(ctx echo.Context) error {
	strId := ctx.Param("id")
	id, _ := strconv.Atoi(strId)

	res, err := gtc.gtu.GetGoodThingById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"errors":err.Error()})	
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"errors":err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

// 作成コントローラー
func (gtc *GoodThingsController) CreatedGoodThing(ctx echo.Context) error {
	goodThing := models.GoodThings{}
	if err := ctx.Bind(&goodThing); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	vError := validation.ValidateStruct(goodThing)
	if vError != nil {
		return ctx.JSON(http.StatusBadRequest, vError)
	}
	res, err := gtc.gtu.CreatedGoodThing(goodThing)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create good thing"})
	}

	return ctx.JSON(http.StatusOK, res)
}

// 更新コントローラー
func (gtc *GoodThingsController) UpdateGoodThing(ctx echo.Context) error {
	goodThing := models.GoodThings{}
	if err := ctx.Bind(&goodThing); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	vError := validation.ValidateStruct(goodThing)
	if vError != nil {
		return ctx.JSON(http.StatusBadRequest, vError)
	}
	strId := ctx.Param("id")
	id, _ := strconv.Atoi(strId)
	res, err := gtc.gtu.UpdateGoodThing(goodThing, id)
	if err != nil {
		if strings.Contains(err.Error(), "レコードが存在しません") {
			return ctx.JSON(http.StatusNotFound, map[string]string{"errors":err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

// 削除コントローラー
func (gtc *GoodThingsController) DeleteGoodThing(ctx echo.Context) error {
	strId := ctx.Param("id")
	id, _ := strconv.Atoi(strId)

	err := gtc.gtu.DeleteGoodThing(id)
	if err != nil {
		if strings.Contains(err.Error(), "レコードが存在しません") {
			return ctx.JSON(http.StatusNotFound, map[string]string{"errors":err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
