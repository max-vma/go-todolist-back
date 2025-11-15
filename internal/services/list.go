package services

import (
	"go-todolist-back/internal/domain"
	"go-todolist-back/internal/domain/repositories"
)

type ListService struct {
	listRepo repositories.ListRepository
}

func NewListService(listRepo repositories.ListRepository) *ListService {
	return &ListService{listRepo: listRepo}
}

func (s *UserService) Register(name, email string) (*domain.User, error) {
	user := &domain.User{Name: name, Email: email}

	return s.userRepo.Create(user)
}

func (s *ListService) GetListsByUserID(id int64) ([]*domain.List, error) {

}
func Create(list *domain.List) error {

}
func Delete(list *domain.List) error {

}
