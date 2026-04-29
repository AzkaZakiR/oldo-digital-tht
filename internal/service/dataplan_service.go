package service

import (
	models "github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	"github.com/AzkaZakiR/oldo-digital-tht/internal/repository"
)

type DataPlanService struct {
	repo repository.DataPlanRepository
}

func NewDataPlanService(repo repository.DataPlanRepository) *DataPlanService {
	return &DataPlanService{repo}
}

func (s *DataPlanService) Create(plan *models.DataPlan) error {
	return s.repo.Create(plan)
}

func (s *DataPlanService) GetAll() ([]models.DataPlan, error) {
	return s.repo.GetAll()
}

func (s *DataPlanService) GetByID(id int) (*models.DataPlan, error) {
	return s.repo.GetByID(id)
}

func (s *DataPlanService) Update(id int, plan *models.DataPlan) error {
	return s.repo.Update(id, plan)
}

func (s *DataPlanService) Delete(id int) error {
	return s.repo.Delete(id)
}