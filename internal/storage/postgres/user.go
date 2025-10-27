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

func (s *UserRepository) Create(user *domain.User) (*domain.User, error) {
	err := s.db.QueryRow(`
	INSERT INTO users (name, email)
	values ($1, $2)
	RETURNING id, name, email, created_at
	`, user.Name, user.Email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
