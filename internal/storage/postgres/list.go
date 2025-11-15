package postgres

import (
	"database/sql"
	"go-todolist-back/internal/domain"
)

type ListRepository struct {
	db *sql.DB
}

func NewListRepository(db *sql.DB) *ListRepository {
	return &ListRepository{db: db}
}

func (s *ListRepository) Create(list *domain.List) (*domain.List, error) {
	err := s.db.QueryRow(`
	INSERT INTO lists (title, user_id)
	values ($1, $2)
	RETURNING id, title, user_id, created_at, updated_at
	`, list.Title, list.User_ID).Scan(&list.ID, &list.Title, &list.User_ID, &list.CreatedAt, &list.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *ListRepository) Delete(list *domain.List) error {
	_, err := s.db.Exec(`
		DELETE FROM lists
		WHERE id = $1
	`, list.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *ListRepository) GetListsByUserID(id int64) ([]domain.List, error) {
	rows, err := s.db.Query(`
	SELECT id, title, user_id, archived, created_at, updated_at 
	FROM lists
	WHERE user_id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []domain.List

	for rows.Next() {
		var l domain.List

		if err := rows.Scan(
			&l.ID,
			&l.Title,
			&l.User_ID,
			&l.Archived,
			&l.CreatedAt,
			&l.UpdatedAt,
		); err != nil {
			return nil, err
		}

		result = append(result, l)
	}

	return result, nil

}
