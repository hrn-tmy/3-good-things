package repository

import (
	"3-good-things/models"

	"gorm.io/gorm"
)

type IGoodThingsRepository interface {
	GetGoodThings(goodThings []*models.GoodThings) error
	GetGoodThingById(goodThing *models.GoodThings, id int) error
	CreatedGoodThing(goodThing *models.GoodThings) error
	UpdateGoodThing(goodThing *models.GoodThings, id int) error
	DeleteGoodThing(id int) error
}

type GoodThingsRepository struct {
	db *gorm.DB
}

func NewGoodThingsRepository(db *gorm.DB) IGoodThingsRepository {
	return &GoodThingsRepository{db}
}

// 一覧取得リポジトリ
func (gtr *GoodThingsRepository) GetGoodThings(goodThings []*models.GoodThings) error {
	if err := gtr.db.Order("id DESC").Find(&goodThings).Error; err != nil {
		return err
	}
	return nil
}

// 詳細取得リポジトリ
func (gtr *GoodThingsRepository) GetGoodThingById(goodThing *models.GoodThings, id int) error {
	if err := gtr.db.Where("id=?", id).First(&goodThing, id).Error; err != nil {
		return err
	}
	return nil
}

// 登録リポジトリ
func (gtr *GoodThingsRepository) CreatedGoodThing(goodThing *models.GoodThings) error {
	if err := gtr.db.Create(&goodThing).Error; err != nil {
		return err
	}
	return nil
}

// 更新リポジトリ
func (gtr *GoodThingsRepository) UpdateGoodThing(goodThing *models.GoodThings, id int) error {
	if err:= gtr.db.Model(&goodThing).Where("id=?", id).Save(&goodThing).Error; err != nil {
		return err
	}
	return nil
}

// 削除リポジトリ
func (gtr *GoodThingsRepository) DeleteGoodThing(id int) error {
	if err := gtr.db.Where("id=?", id).Delete(&models.GoodThings{}).Error; err != nil {
		return err
	}
	return nil
}