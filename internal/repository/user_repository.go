package repository

import (
	models "github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface{
	GetAll() ([]models.User, error)
	GetById(id int) (*models.User, error)
	Create(user *models.User) error
	Update(id int, user *models.User) error
	Delete(id int) error
}

type userRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userRepository{db: db}
}
func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetById(id int) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(id int, user *models.User) error {
		return r.db.Model(&models.User{}).
		Where("id = ?", id).Updates(user).Error
}

func (r *userRepository) Delete(id int) error{
	return r.db.Delete(&models.User{}, id).Error
}
