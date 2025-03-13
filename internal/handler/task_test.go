package handler

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gofiber/fiber/v2"
    "github.com/guttof/task-api-go/internal/domain"
    "github.com/guttof/task-api-go/internal/repository"
    "github.com/guttof/task-api-go/internal/service"
)

func setupTest() (*fiber.App, *TaskHandler) {
    app := fiber.New()
    repo := repository.NewInMemoryTaskRepository()
    svc := service.NewTaskService(repo)
    handler := NewTaskHandler(svc)
    return app, handler
}

func TestCreateTask(t *testing.T) {
    app, handler := setupTest()
    app.Post("/tasks", handler.CreateTask)

    task := domain.Task{ID: "1", Title: "Test Task"}
    body, _ := json.Marshal(task)

    req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    resp, err := app.Test(req)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if resp.StatusCode != http.StatusCreated {
        t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
    }
}
