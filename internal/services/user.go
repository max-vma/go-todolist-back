package services

import (
	"go-todolist-back/internal/domain"
	"go-todolist-back/internal/domain/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(name, email string) (*domain.User, error) {
	user := &domain.User{Name: name, Email: email}

	return s.userRepo.Create(user)
}
