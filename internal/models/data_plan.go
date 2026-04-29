package models

import "gorm.io/gorm"

type DataPlan struct{
	gorm.Model
	Name string `json:"name"`
	Price int `json:"price"`
	Quota int `json:"quota"`
	ActivePeriod int `json:"active_period"`
	IsActive bool `json:"is_active"`

	Transactions []Transaction `gorm:"foreignKey:DataPlanID"`
}