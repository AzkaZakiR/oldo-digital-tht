package dto

import "github.com/gofiber/fiber/v3"


type APIResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}
func Success(c fiber.Ctx, data interface{}, message string) error {
	return c.JSON(APIResponse{
		Status:  "success",
		Data:    data,
		Message: message,
		Errors:  nil,
	})
}

func Error(c fiber.Ctx, code int, message string, errors interface{}) error {
	return c.Status(code).JSON(APIResponse{
		Status:  "error",
		Data:    nil,
		Message: message,
		Errors:  errors,
	})
}