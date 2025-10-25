package postgres

import (
	"database/sql"
	"go-todolist-back/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (s *UserRepository) Create(user *domain.User) error {
	return nil
}
