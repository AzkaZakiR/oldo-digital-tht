package main

import (
	"log"

	"github.com/AzkaZakiR/oldo-digital-tht/internal/database"
	"github.com/AzkaZakiR/oldo-digital-tht/internal/handler"
	models "github.com/AzkaZakiR/oldo-digital-tht/internal/models"
	dto "github.com/AzkaZakiR/oldo-digital-tht/internal/pkg"
	"github.com/AzkaZakiR/oldo-digital-tht/internal/repository"
	"github.com/AzkaZakiR/oldo-digital-tht/internal/service"
	"github.com/gofiber/fiber/v3"
)
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
func main(){
	db, err := database.OpenConnection()
	if err != nil{
		log.Fatal(err)
	}
	err = database.Migrate(db)
	if err != nil{
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)
	dataPlanRepo := repository.NewDataPlanRepository(db)
	dataPlanSvc := service.NewDataPlanService(dataPlanRepo)
	dataPlanHandler := handler.NewDataPlanHandler(dataPlanSvc)
	transactionRepo := repository.NewTransactionRepository(db)
	transactionSvc := service.NewTransactionService(transactionRepo, dataPlanRepo, userRepo)
	transactionHandler := handler.NewTransactionHandler(transactionSvc)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
        code := fiber.StatusInternalServerError
        if e, ok := err.(*fiber.Error); ok {
            code = e.Code
        }
        return dto.Error(c, code, "An unexpected error occurred", err.Error())
    },
	})

	userApi := app.Group("/api/users")
	dataPlanApi := app.Group("/api/dataplan")
	transactionApi := app.Group("/api/transactions")

	userApi.Get("", userHandler.GetAll)
	userApi.Get("/:id", userHandler.GetByID)
	userApi.Post("", userHandler.Create)
	userApi.Patch("/:id", userHandler.Update)
	userApi.Delete("/:id", userHandler.Delete)

	dataPlanApi.Get("", dataPlanHandler.GetAll)
	dataPlanApi.Get("/:id", dataPlanHandler.GetByID)
	dataPlanApi.Post("", dataPlanHandler.Create)
	dataPlanApi.Patch("/:id", dataPlanHandler.Update)
	dataPlanApi.Delete("/:id", dataPlanHandler.Delete)


	transactionApi.Get("", transactionHandler.GetAll)
	transactionApi.Get("/:id", transactionHandler.GetByID)
	transactionApi.Post("", transactionHandler.Create)

	app.Get("/", func (c fiber.Ctx) error  {
		res := Response{
			Message: "Hello from Gofibber",
			Status:  200,
		}
		return c.JSON(res)
	})

	app.Put("/users/:id", func (c fiber.Ctx) error  {
		id := c.Params("id")

		var user models.User

		if err := db.First(&user, id).Error; err != nil{
			return c.Status(404).JSON("User not found")
		}

		var updateData models.User
		if err := c.Bind().Body(&updateData); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		db.Model(&user).Updates(updateData)

		return c.JSON(user)
	})

	app.Listen(":3000")
}