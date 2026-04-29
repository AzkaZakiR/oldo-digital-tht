package repository

import (
	models "github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAll() ([]models.Transaction, error)
	GetByID(id int) (*models.Transaction, error)
	Create(tx *models.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetAll() ([]models.Transaction, error) {
	var txs []models.Transaction
	err := r.db.
		Preload("User").
		Preload("DataPlan").
		Find(&txs).Error

	return txs, err
}

func (r *transactionRepository) GetByID(id int) (*models.Transaction, error) {
	var tx models.Transaction
	err := r.db.
		Preload("User").
		Preload("DataPlan").
		First(&tx, id).Error

	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func (r *transactionRepository) Create(tx *models.Transaction) error {
	return r.db.Create(tx).Error
}