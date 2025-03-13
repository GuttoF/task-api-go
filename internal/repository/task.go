package repository

import (
    "errors"
    "github.com/guttof/task-api-go/internal/domain"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository interface {
    Create(task domain.Task) error
    FindAll() ([]domain.Task, error)
    FindByID(id string) (*domain.Task, error)
    Update(task domain.Task) error
    Delete(id string) error
}

type inMemoryTaskRepository struct {
    tasks []domain.Task
}

func NewInMemoryTaskRepository() TaskRepository {
    return &inMemoryTaskRepository{
        tasks: make([]domain.Task, 0),
    }
}

func (r *inMemoryTaskRepository) Create(task domain.Task) error {
    r.tasks = append(r.tasks, task)
    return nil
}

func (r *inMemoryTaskRepository) FindAll() ([]domain.Task, error) {
    return r.tasks, nil
}

func (r *inMemoryTaskRepository) FindByID(id string) (*domain.Task, error) {
    for _, task := range r.tasks {
        if task.ID == id {
            return &task, nil
        }
    }
    return nil, ErrTaskNotFound
}

func (r *inMemoryTaskRepository) Update(task domain.Task) error {
    for i, t := range r.tasks {
        if t.ID == task.ID {
            r.tasks[i] = task
            return nil
        }
    }
    return ErrTaskNotFound
}

func (r *inMemoryTaskRepository) Delete(id string) error {
    for i, task := range r.tasks {
        if task.ID == id {
            r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
            return nil
        }
    }
    return ErrTaskNotFound
}
