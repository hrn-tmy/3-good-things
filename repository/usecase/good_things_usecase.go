package usecase

import (
	"3-good-things/models"
	"3-good-things/repository"
)

type IGoodThingsUsecase interface {
	GetGoodThings() ([]models.GoodThingsResponse, error)
	GetGoodThingById(id int) (models.GoodThingsResponse, error)
	CreatedGoodThing(goodThing models.GoodThings) (models.GoodThingsResponse, error)
	UpdateGoodThing(goodThing models.GoodThings, id int) (models.GoodThingsResponse, error)
	DeleteGoodThing(id int) error
}

type GoodThingsUsecase struct {
	gtr repository.IGoodThingsRepository
}

func NewGoodThingsUsecase(gtr repository.IGoodThingsRepository) IGoodThingsUsecase {
	return &GoodThingsUsecase{gtr}
}

// 一覧ユースケース
func (gtu *GoodThingsUsecase) GetGoodThings() ([]models.GoodThingsResponse, error) {
	goodThings := []models.GoodThings{}
	if err := gtu.gtr.GetGoodThings(&goodThings); err != nil {
		return []models.GoodThingsResponse{}, err
	}
	res := []models.GoodThingsResponse{}
	for _, v := range goodThings {
		g := models.GoodThingsResponse{
			ID:                v.ID,
			Date:              v.Date,
			FirstGoodThing:    v.FirstGoodThing,
			SecondGoodThing:   v.SecondGoodThing,
			ThirdGoodThing:    v.ThirdGoodThing,
			FirstImprovement:  v.FirstImprovement,
			SecondImprovement: v.SecondImprovement,
			ThirdImprovement:  v.ThirdImprovement,
			FristLearntThing:  v.FristLearntThing,
			SecondLearntThing: v.SecondLearntThing,
			ThirdLearntThing:  v.ThirdLearntThing,
		}
		res = append(res, g)
	}
	return res, nil
}

// 詳細取得ユースケース
func (gtu *GoodThingsUsecase) GetGoodThingById(id int) (models.GoodThingsResponse, error) {
	goodThing := models.GoodThings{}
	if err := gtu.gtr.GetGoodThingById(&goodThing, id); err != nil {
		return models.GoodThingsResponse{}, err
	}
	res := models.GoodThingsResponse{
		ID:                goodThing.ID,
		Date:              goodThing.Date,
		FirstGoodThing:    goodThing.FirstGoodThing,
		SecondGoodThing:   goodThing.SecondGoodThing,
		ThirdGoodThing:    goodThing.ThirdGoodThing,
		FirstImprovement:  goodThing.FirstImprovement,
		SecondImprovement: goodThing.SecondImprovement,
		ThirdImprovement:  goodThing.ThirdImprovement,
		FristLearntThing:  goodThing.FristLearntThing,
		SecondLearntThing: goodThing.SecondLearntThing,
		ThirdLearntThing:  goodThing.ThirdLearntThing,
	}
	return res, nil
}

// 作成ユースケース
func (gtu *GoodThingsUsecase) CreatedGoodThing(goodThing models.GoodThings) (models.GoodThingsResponse, error) {
	if err := gtu.gtr.CreatedGoodThing(&goodThing); err != nil {
		return models.GoodThingsResponse{}, err
	}
	res := models.GoodThingsResponse{
		ID:                goodThing.ID,
		Date:              goodThing.Date,
		FirstGoodThing:    goodThing.FirstGoodThing,
		SecondGoodThing:   goodThing.SecondGoodThing,
		ThirdGoodThing:    goodThing.ThirdGoodThing,
		FirstImprovement:  goodThing.FirstImprovement,
		SecondImprovement: goodThing.SecondImprovement,
		ThirdImprovement:  goodThing.ThirdImprovement,
		FristLearntThing:  goodThing.FristLearntThing,
		SecondLearntThing: goodThing.SecondLearntThing,
		ThirdLearntThing:  goodThing.ThirdLearntThing,
	}
	return res, nil
}

// 更新ユースケース
func (gtu *GoodThingsUsecase) UpdateGoodThing(goodThing models.GoodThings, id int) (models.GoodThingsResponse, error) {
	if err := gtu.gtr.UpdateGoodThing(&goodThing, id); err != nil {
		return models.GoodThingsResponse{}, err
	}
	res := models.GoodThingsResponse{
		ID:                goodThing.ID,
		Date:              goodThing.Date,
		FirstGoodThing:    goodThing.FirstGoodThing,
		SecondGoodThing:   goodThing.SecondGoodThing,
		ThirdGoodThing:    goodThing.ThirdGoodThing,
		FirstImprovement:  goodThing.FirstImprovement,
		SecondImprovement: goodThing.SecondImprovement,
		ThirdImprovement:  goodThing.ThirdImprovement,
		FristLearntThing:  goodThing.FristLearntThing,
		SecondLearntThing: goodThing.SecondLearntThing,
		ThirdLearntThing:  goodThing.ThirdLearntThing,
	}
	return res, nil
}

// 削除ユースケース
func (gtu *GoodThingsUsecase) DeleteGoodThing(id int) error {
	if err := gtu.gtr.DeleteGoodThing(id); err != nil {
		return err
	}
	return nil
}
