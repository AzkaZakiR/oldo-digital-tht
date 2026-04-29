package service

import (
	"errors"

	models "github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	"github.com/AzkaZakiR/oldo-digital-tht/internal/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
	dataPlanRepo repository.DataPlanRepository
	userRepo repository.UserRepository
}

func NewTransactionService(
	repo repository.TransactionRepository,
	dataPlanRepo repository.DataPlanRepository,
	userRepo repository.UserRepository,
) *TransactionService {
	return &TransactionService{repo, dataPlanRepo, userRepo}
}

func (s *TransactionService) Create(tx *models.Transaction) (*models.Transaction, error) {
	plan, err := s.dataPlanRepo.GetByID(tx.DataPlanID)
	if err != nil {
		return nil, errors.New("Data plan not found")
	}
	_, err = s.userRepo.GetById(tx.UserID)
	if err != nil {
		return nil, errors.New("User not found")
	}

	if !plan.IsActive{
		return nil, errors.New("data plan is not active")
	}
	tx.Price = plan.Price

	if err := s.repo.Create(tx); err != nil {
		return nil, err
	}
	return  s.repo.GetByID(int(tx.ID))
}

func (s *TransactionService) GetAll() ([]models.Transaction, error) {
	return s.repo.GetAll()
}

func (s *TransactionService) GetByID(id int) (*models.Transaction, error) {
	return s.repo.GetByID(id)
}