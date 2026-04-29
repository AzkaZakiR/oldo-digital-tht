package models

import "gorm.io/gorm"

type Transaction struct{
	gorm.Model
	UserID     int `json:"user_id"`
	DataPlanID int `json:"data_plan_id"`
	Price int `json:"price"`

	User User `gorm:"foreignKey:UserID"`
	DataPlan DataPlan `gorm:"foreignKey:DataPlanID"`
}