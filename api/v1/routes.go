package v1

import (
    "github.com/gofiber/fiber/v2"
    "github.com/guttof/task-api-go/internal/handler"
)

func SetupRoutes(app *fiber.App, taskHandler *handler.TaskHandler) {
    v1 := app.Group("/api/v1")
    
    tasks := v1.Group("/tasks")
    tasks.Post("/", taskHandler.CreateTask)
    tasks.Get("/", taskHandler.GetTasks)
    tasks.Get("/:id", taskHandler.GetTask)
    tasks.Put("/:id", taskHandler.UpdateTask)
    tasks.Delete("/:id", taskHandler.DeleteTask)
}
