package handler

import (
	"strconv"

	"github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	dto "github.com/AzkaZakiR/oldo-digital-tht/internal/pkg"
	"github.com/AzkaZakiR/oldo-digital-tht/internal/service"
	"github.com/gofiber/fiber/v3"
)

type TransactionHandler struct {
	svc *service.TransactionService
}

func NewTransactionHandler(svc *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{svc}
}

type CreateTransactionRequest struct {
	UserID     int `json:"user_id" validate:"required"`
	DataPlanID int `json:"data_plan_id" validate:"required"`
}

func (h *TransactionHandler) GetAll(c fiber.Ctx) error {
	txs, err := h.svc.GetAll()
	if err != nil {
		return dto.Error(c, 500, "failed to get transactions", err.Error())
	}
	return dto.Success(c, txs, "transactions retrieved")
}

func (h *TransactionHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return dto.Error(c, 400, "invalid id", err.Error())
	}

	tx, err := h.svc.GetByID(id)
	if err != nil {
		return dto.Error(c, 404, "transaction not found", err.Error())
	}

	return dto.Success(c, tx, "transaction retrieved")
}

func (h *TransactionHandler) Create(c fiber.Ctx) error {
	var req CreateTransactionRequest

	if err := c.Bind().Body(&req); err != nil {
		return dto.Error(c, 400, "invalid request", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return dto.Error(c, 400, "validation error", err.Error())
	}


	tx, err := h.svc.Create(&models.Transaction{
	UserID:     req.UserID,
	DataPlanID: req.DataPlanID,
	})

	if err != nil {
		return dto.Error(c, 500, "failed to create transaction", err.Error())
	}

	return dto.Success(c, tx, "transaction created")
}