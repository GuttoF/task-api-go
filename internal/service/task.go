package service

import (
    "github.com/guttof/task-api-go/internal/domain"
    "github.com/guttof/task-api-go/internal/repository"
)

type TaskService interface {
    Create(task domain.Task) error
    GetAll() ([]domain.Task, error)
    GetByID(id string) (*domain.Task, error)
    Update(task domain.Task) error
    Delete(id string) error
}

type taskService struct {
    repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
    return &taskService{
        repo: repo,
    }
}

func (s *taskService) Create(task domain.Task) error {
    return s.repo.Create(task)
}

func (s *taskService) GetAll() ([]domain.Task, error) {
    return s.repo.FindAll()
}

func (s *taskService) GetByID(id string) (*domain.Task, error) {
    return s.repo.FindByID(id)
}

func (s *taskService) Update(task domain.Task) error {
    return s.repo.Update(task)
}

func (s *taskService) Delete(id string) error {
    return s.repo.Delete(id)
}
