package handler

import (
    "github.com/gofiber/fiber/v2"
    "github.com/guttof/task-api-go/internal/domain"
    "github.com/guttof/task-api-go/internal/service"
)

type TaskHandler struct {
    svc service.TaskService
}

func NewTaskHandler(svc service.TaskService) *TaskHandler {
    return &TaskHandler{
        svc: svc,
    }
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
    var task domain.Task
    if err := c.BodyParser(&task); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := h.svc.Create(task); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
    tasks, err := h.svc.GetAll()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(tasks)
}

func (h *TaskHandler) GetTask(c *fiber.Ctx) error {
    id := c.Params("id")
    task, err := h.svc.GetByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Task not found",
        })
    }

    return c.JSON(task)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
    var task domain.Task
    if err := c.BodyParser(&task); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    task.ID = c.Params("id")
    if err := h.svc.Update(task); err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Task not found",
        })
    }

    return c.JSON(task)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
    id := c.Params("id")
    if err := h.svc.Delete(id); err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Task not found",
        })
    }

    return c.SendStatus(fiber.StatusNoContent)
}
