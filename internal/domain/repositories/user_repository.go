package repositories

import "go-todolist-back/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
}
