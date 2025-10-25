package repositories

import (
	"go-todolist-back/internal/domain"
)

type TodoRepository interface {
	Create(todo *domain.Todo) error
	GetByID(id int64) (*domain.Todo, error)
	DeleteByID(id int64) error
	Delete(todo *domain.Todo) error
	UpdateById(todo domain.Todo) error
}
