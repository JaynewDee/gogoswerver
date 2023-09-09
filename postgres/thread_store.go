package postgres

import (
	"fmt"

	"github.com/gogoswerver"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) Thread(id uuid.UUID) (gogoswerver.Thread, error) {
	var t gogoswerver.Thread

	if err := s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		return gogoswerver.Thread{}, fmt.Errorf("error getting threads: %w", err)

	}

	return t, nil
}

func (s *ThreadStore) Threads() ([]gogoswerver.Thread, error) {
	var ts []gogoswerver.Thread

	if err := s.Select(&ts, `SELECT * FROM threads`); err != nil {
		return []gogoswerver.Thread{}, fmt.Errorf("error getting threads: %w", err)
	}

	return ts, nil
}

func (s *ThreadStore) CreateThread(t *gogoswerver.Thread) error {
	if err := s.Get(t, `INSERT INTO threads VALUES ($1, $2, $3) RETURNING *`,
		t.ID,
		t.Title,
		t.Description,
	); err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}

	return nil
}

func (s *ThreadStore) UpdateThread(t *gogoswerver.Thread) error {
	if err := s.Get(t, `UPDATE threads SET title = $1, description = $2 WHERE id = $3`,
		t.Title,
		t.Description,
		t.ID,
	); err != nil {
		return fmt.Errorf("error updating thread: %w", err)
	}

	return nil
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting thread: %w", err)
	}

	return nil
}
