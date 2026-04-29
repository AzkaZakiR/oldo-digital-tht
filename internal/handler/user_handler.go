package handler

import (
	"strconv"

	"github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	dto "github.com/AzkaZakiR/oldo-digital-tht/internal/pkg"
	"github.com/AzkaZakiR/oldo-digital-tht/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
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
		return dto.Error(c, 400, "error creating user", err.Error())
	}
	if err := validate.Struct(req); err != nil {
		return dto.Error(c, 400, "please use the correct format", err.Error())
	}

	user := models.User{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}
	if err := h.svc.CreateUser(&user); err != nil {
		return dto.Error(c, 500, "error creating user", err.Error())
	}
	return dto.Success(c, user, "user created")
}

func (h *UserHandler) GetAll(c fiber.Ctx) error {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		return dto.Error(c, 500, "error fetching users", err.Error())
	}
	return dto.Success(c, users, "users fetched")
}

func (h *UserHandler) GetByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return dto.Error(c, 400, "invalid id", err.Error())
	}
	user, err := h.svc.GetUserByID(int(id))
	if err != nil {
		return dto.Error(c, 404, "user not found", err.Error())
	}
	return dto.Success(c, user, "user fetched")
}

func (h *UserHandler) Update(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return dto.Error(c, 400, "invalid id", err.Error())
	}

	var req UpdateUserRequest
	if err := c.Bind().Body(&req); err != nil {
		return dto.Error(c, 400, "error updating user", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return dto.Error(c, 400, "please use the correct format", err.Error())
	}

	user, err := h.svc.GetUserByID(id)
	if err != nil {
		return dto.Error(c, 404, "user not found", err.Error())
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
		return dto.Error(c, 500, "error updating user", err.Error())
	}

	return dto.Success(c, user, "user updated")
}

func (h *UserHandler) Delete(c fiber.Ctx) error {
idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return dto.Error(c, 400, "invalid id", err.Error())
	}
	if err := h.svc.DeleteUser(int(id)); err != nil {
		return dto.Error(c, 500, "error deleting user", err.Error())
	}
	return dto.Success(c, nil, "user deleted")
}