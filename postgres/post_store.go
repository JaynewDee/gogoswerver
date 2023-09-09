package postgres

import (
	"fmt"

	"github.com/JaynewDee/gogoswerver"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (gogoswerver.Post, error) {
	var p gogoswerver.Post

	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return gogoswerver.Post{}, fmt.Errorf("error retrieving post by id: %w", err)
	}

	return p, nil
}

func (s *PostStore) PostsByThread(threadID uuid.UUID) ([]gogoswerver.Post, error) {
	var ps []gogoswerver.Post

	if err := s.Select(&ps, `SELECT * FROM posts WHERE thread_id = $1`, threadID); err != nil {
		return []gogoswerver.Post{}, err
	}

	return ps, nil
}

func (s *PostStore) CreatePost(p *gogoswerver.Post) error {
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

func (s *PostStore) UpdatePost(p *gogoswerver.Post) error {
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
