// internal/repository/task_test.go
package repository

import (
    "testing"
    "github.com/guttof/task-api-go/internal/domain"
)

func TestCreateTask(t *testing.T) {
    repo := NewInMemoryTaskRepository()
    task := domain.Task{ID: "1", Title: "Test Task"}
    
    err := repo.Create(task)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    
    tasks, _ := repo.FindAll()
    if len(tasks) != 1 {
        t.Fatalf("Expected 1 task, got %d", len(tasks))
    }
}
