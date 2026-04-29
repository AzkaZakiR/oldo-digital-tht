package repository

import (
	models "github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	"gorm.io/gorm"
)

type DataPlanRepository interface {
	GetAll() ([]models.DataPlan, error)
	GetByID(id int) (*models.DataPlan, error)
	Create(dataPlan *models.DataPlan) error
	Update(id int, dataPlan *models.DataPlan) error
	Delete(id int) error
}

type dataPlanRepository struct {
	db *gorm.DB
}

func NewDataPlanRepository(db *gorm.DB) DataPlanRepository {
	return &dataPlanRepository{db}
}

func (r *dataPlanRepository) GetAll() ([]models.DataPlan, error) {
	var plans []models.DataPlan
	err := r.db.Find(&plans).Error
	return plans, err
}

func (r *dataPlanRepository) GetByID(id int) (*models.DataPlan, error) {
	var plan models.DataPlan
	err := r.db.First(&plan, id).Error
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

func (r *dataPlanRepository) Create(dataPlan *models.DataPlan) error {
	return r.db.Create(dataPlan).Error
}

func (r *dataPlanRepository) Update(id int, dataPlan *models.DataPlan) error {
	return r.db.Model(&models.DataPlan{}).
		Where("id = ?", id).
		Updates(dataPlan).Error
}

func (r *dataPlanRepository) Delete(id int) error {
	return r.db.Delete(&models.DataPlan{}, id).Error
}