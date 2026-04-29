package handler

import (
	"strconv"

	"github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	dto "github.com/AzkaZakiR/oldo-digital-tht/internal/pkg"
	"github.com/AzkaZakiR/oldo-digital-tht/internal/service"
	"github.com/gofiber/fiber/v3"
)

type DataPlanHandler struct {
	svc *service.DataPlanService
}

func NewDataPlanHandler(svc *service.DataPlanService) *DataPlanHandler {
	return &DataPlanHandler{svc}
}

type CreateDataPlanRequest struct {
	Name         string `json:"name" validate:"required"`
	Price        int    `json:"price" validate:"required"`
	Quota        int    `json:"quota" validate:"required"`
	ActivePeriod int    `json:"active_period" validate:"required"`
	IsActive     bool   `json:"is_active"`
}

func (h *DataPlanHandler) GetAll(c fiber.Ctx) error {
	plans, err := h.svc.GetAll()
	if err != nil {
		return dto.Error(c, 500, "failed to get data plans", err.Error())
	}
	return dto.Success(c, plans, "data plans retrieved")
}

func (h *DataPlanHandler) GetByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return dto.Error(c, 400, "invalid id", err.Error())
	}

	plan, err := h.svc.GetByID(int(id))
	if err != nil {
		return dto.Error(c, 404, "data plan not found", err.Error())
	}

	return dto.Success(c, plan, "data plan retrieved")
}


func (h *DataPlanHandler) Create(c fiber.Ctx) error {
	var req CreateDataPlanRequest

	if err := c.Bind().Body(&req); err != nil {
		return dto.Error(c, 400, "invalid request", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return dto.Error(c, 400, "validation error", err.Error())
	}

	plan := models.DataPlan{
		Name:         req.Name,
		Price:        req.Price,
		Quota:        req.Quota,
		ActivePeriod: req.ActivePeriod,
		IsActive:     req.IsActive,
	}

	if err := h.svc.Create(&plan); err != nil {
		return dto.Error(c, 500, "failed to create data plan", err.Error())
	}

	return dto.Success(c, plan, "data plan created")
}

func (h *DataPlanHandler) Update(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return dto.Error(c, 400, "invalid id", err.Error())
	}

	var req CreateDataPlanRequest
	if err := c.Bind().Body(&req); err != nil {
		return dto.Error(c, 400, "invalid request", err.Error())
	}

	plan, err := h.svc.GetByID(int(id))
	if err != nil {
		return dto.Error(c, 404, "data plan not found", err.Error())
	}

	plan.Name = req.Name
	plan.Price = req.Price
	plan.Quota = req.Quota
	plan.ActivePeriod = req.ActivePeriod
	plan.IsActive = req.IsActive

	if err := h.svc.Update(int(id), plan); err != nil {
		return dto.Error(c, 500, "failed to update", err.Error())
	}

	return dto.Success(c, plan, "data plan updated")
}

func (h *DataPlanHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return dto.Error(c, 400, "invalid id", err.Error())
	}

	if err := h.svc.Delete(int(id)); err != nil {
		return dto.Error(c, 500, "failed to delete", err.Error())
	}

	return dto.Success(c, nil, "data plan deleted")
}