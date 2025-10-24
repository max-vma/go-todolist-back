package repositories

import "go-todolist-back/internal/domain"

type ListRepository interface {
	GetListsByUserID(id int64) ([]*domain.List, error)
	Create(list *domain.List) error
	DeleteByID(id int64) error
	Delete(list *domain.List) error
}
