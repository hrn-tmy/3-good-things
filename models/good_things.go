package models

import "time"

type GoodThings struct {
	ID                int       `json:"id" gorm:"PrimaryKey"`
	Date              time.Time `json:"date"`
	FirstGoodThing    string    `json:"first_good_thing" validate:"required,min=1,max=255"`
	SecondGoodThing   string    `json:"second_good_thing" validate:"required,min=1,max=255"`
	ThirdGoodThing    string    `json:"third_good_thing" validate:"required,min=1,max=255"`
	FirstImprovement  string    `json:"first_improvement" validate:"required,min=1,max=255"`
	SecondImprovement string    `json:"second_improvement" validate:"required,min=1,max=255"`
	ThirdImprovement  string    `json:"third_improvement" validate:"required,min=1,max=255"`
	FristLearntThing  string    `json:"first_learnt_thing" validate:"required,min=1,max=255"`
	SecondLearntThing string    `json:"second_learnt_thing" validate:"required,min=1,max=255"`
	ThirdLearntThing  string    `json:"third_learnt_thing" validate:"required,min=1,max=255"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
}

type GoodThingsResponse struct {
	ID                int       `json:"id" gorm:"PrimaryKey"`
	Date              time.Time `json:"date"`
	FirstGoodThing    string    `json:"first_good_thing"`
	SecondGoodThing   string    `json:"second_good_thing"`
	ThirdGoodThing    string    `json:"third_good_thing"`
	FirstImprovement  string    `json:"first_improvement"`
	SecondImprovement string    `json:"second_improvement"`
	ThirdImprovement  string    `json:"third_improvement"`
	FristLearntThing  string    `json:"first_learnt_thing"`
	SecondLearntThing string    `json:"second_learnt_thing"`
	ThirdLearntThing  string    `json:"third_learnt_thing"`
}
