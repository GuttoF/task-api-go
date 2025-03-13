package service

import (
    "testing"
    "github.com/guttof/task-api-go/internal/domain"
    "github.com/guttof/task-api-go/internal/repository"
)

func TestTaskService_Create(t *testing.T) {
    repo := repository.NewInMemoryTaskRepository()
    service := NewTaskService(repo)
    
    task := domain.Task{ID: "1", Title: "Test Task"}
    
    err := service.Create(task)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
}

func TestTaskService_GetAll(t *testing.T) {
    repo := repository.NewInMemoryTaskRepository()
    service := NewTaskService(repo)
    
    task1 := domain.Task{ID: "1", Title: "Task 1"}
    task2 := domain.Task{ID: "2", Title: "Task 2"}
    
    service.Create(task1)
    service.Create(task2)
    
    tasks, err := service.GetAll()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    
    if len(tasks) != 2 {
        t.Fatalf("Expected 2 tasks, got %d", len(tasks))
    }
}
