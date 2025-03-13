package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/guttof/task-api-go/api/v1"
    "github.com/guttof/task-api-go/internal/handler"
    "github.com/guttof/task-api-go/internal/repository"
    "github.com/guttof/task-api-go/internal/service"
)

func main() {
    app := fiber.New(fiber.Config{
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        },
    })

    app.Use(logger.New())

    // Setup dependencies
    taskRepo := repository.NewInMemoryTaskRepository()
    taskService := service.NewTaskService(taskRepo)
    taskHandler := handler.NewTaskHandler(taskService)

    // Setup routes
    v1.SetupRoutes(app, taskHandler)

    log.Fatal(app.Listen(":8080"))
}
