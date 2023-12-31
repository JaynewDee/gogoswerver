package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jaynewdee/gogoswerver/entity"
	"github.com/jmoiron/sqlx"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (entity.Post, error) {
	var p entity.Post

	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return entity.Post{}, fmt.Errorf("error retrieving post by id: %w", err)
	}

	return p, nil
}

func (s *PostStore) PostsByThread(threadID uuid.UUID) ([]entity.Post, error) {
	var ps []entity.Post

	if err := s.Select(&ps, `SELECT * FROM posts WHERE thread_id = $1`, threadID); err != nil {
		return []entity.Post{}, err
	}

	return ps, nil
}

func (s *PostStore) CreatePost(p *entity.Post) error {
	if err := s.Get(&p, `INSERT INTO posts VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		p.ID,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes,
	); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}

	return nil
}

func (s *PostStore) UpdatePost(p *entity.Post) error {
	if err := s.Get(&p, `UPDATE posts SET thread_id = $1, title = $2, content = $3, votes = $4 WHERE id = $5`,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes,
		p.ID,
	); err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}

	return nil
}

func (s *PostStore) DeletePost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}

	return nil
}
