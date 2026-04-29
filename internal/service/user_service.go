package service

import (
	"errors"

	models "github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	"github.com/AzkaZakiR/oldo-digital-tht/internal/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailExists = errors.New("user with this email already exists")
type UserService struct{
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}
func (s *UserService) CreateUser(user *models.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	
	err = s.repo.Create(user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrEmailExists
		}
		return err
	}
	return nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) UpdateUser(id int, user *models.User) error {
	if user.Password != "" {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	return s.repo.Update(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}