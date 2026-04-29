package handler

import (
	"errors"
	"strconv"

	"github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	dto "github.com/AzkaZakiR/oldo-digital-tht/internal/pkg"
	"github.com/AzkaZakiR/oldo-digital-tht/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)
type UserHandler struct {
	svc *service.UserService
}

type CreateUserRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required,numeric"`
	Password    string `json:"password" validate:"required,min=6"`
}
type UpdateUserRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=3"`
	Email       *string `json:"email" validate:"omitempty,email"`
	PhoneNumber *string `json:"phoneNumber" validate:"omitempty,numeric"`
	Password    *string `json:"password" validate:"omitempty,min=6"`
}
var validate = validator.New() 

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc}
}

func (h *UserHandler) Create(c fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.Bind().Body(&req); err != nil {
		return dto.Error(c, 400, "error creating user", dto.FormatValidationError(err))
	}
	if err := validate.Struct(req); err != nil {
		return dto.Error(c, 400, "please use the correct format", dto.FormatValidationError(err))
	}

	user := models.User{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}
	if err := h.svc.CreateUser(&user); err != nil {
		if errors.Is(err, service.ErrEmailExists) {
            return dto.Error(c, 409, "Registration failed", map[string]string{
                "email": "email already in use",
            })
        }
		return dto.Error(c, 500, "error creating user", dto.FormatValidationError(err))
	}
	return dto.Success(c, user, "user created")
}

func (h *UserHandler) GetAll(c fiber.Ctx) error {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		return dto.Error(c, 500, "error fetching users", dto.FormatValidationError(err))
	}
	return dto.Success(c, users, "users fetched")
}

func (h *UserHandler) GetByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return dto.Error(c, 400, "invalid id", dto.FormatValidationError(err))
	}
	user, err := h.svc.GetUserByID(int(id))
	if err != nil {
		return dto.Error(c, 404, "user not found", dto.FormatValidationError(err))
	}
	return dto.Success(c, user, "user fetched")
}

func (h *UserHandler) Update(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return dto.Error(c, 400, "invalid id", dto.FormatValidationError(err))
	}

	var req UpdateUserRequest
	if err := c.Bind().Body(&req); err != nil {
		return dto.Error(c, 400, "error updating user", dto.FormatValidationError(err))
	}

	if err := validate.Struct(req); err != nil {
		return dto.Error(c, 400, "please use the correct format", dto.FormatValidationError(err))
	}

	user, err := h.svc.GetUserByID(id)
	if err != nil {
		return dto.Error(c, 404, "user not found", dto.FormatValidationError(err))
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.PhoneNumber != nil {
		user.PhoneNumber = *req.PhoneNumber
	}
	if req.Password != nil {
		user.Password = *req.Password
	}

	if err := h.svc.UpdateUser(id, user); err != nil {
		return dto.Error(c, 500, "error updating user", dto.FormatValidationError(err))
	}

	return dto.Success(c, user, "user updated")
}

func (h *UserHandler) Delete(c fiber.Ctx) error {
idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return dto.Error(c, 400, "invalid id", dto.FormatValidationError(err))
	}
	if err := h.svc.DeleteUser(int(id)); err != nil {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.Error(c, 404, "user not found", nil)
	}
		return dto.Error(c, 500, "error deleting user", err.Error())
	}
	return dto.Success(c, nil, "user deleted")
}