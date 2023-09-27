package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jaynewdee/gogoswerver/entity"
	"github.com/jmoiron/sqlx"
)

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id uuid.UUID) (entity.Comment, error) {
	var c entity.Comment

	if err := s.Get(&c, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return entity.Comment{}, fmt.Errorf("error retrieving comment: %w", err)
	}

	return c, nil
}

func (s *CommentStore) CommentsByPost(postID uuid.UUID) ([]entity.Comment, error) {
	var cs []entity.Comment

	if err := s.Select(&cs, `SELECT * FROM comments WHERE post_id = $1`, postID); err != nil {
		return []entity.Comment{}, err
	}

	return cs, nil
}

func (s *CommentStore) CreateComment(c *entity.Comment) error {
	if err := s.Get(&c, `INSERT INTO comments VALUES ($1, $2, $3, $4) RETURNING *`,
		c.ID,
		c.PostID,
		c.Content,
		c.Votes,
	); err != nil {
		return fmt.Errorf("error creating comment: %w", err)
	}

	return nil
}

func (s *CommentStore) UpdateComment(c *entity.Comment) error {
	if err := s.Get(&c, `UPDATE comments SET post_id = $1, content = $2, votes = $3 WHERE post_id = $4`,
		c.PostID,
		c.Content,
		c.Votes,
		c.ID,
	); err != nil {
		return fmt.Errorf("error updating comment: %w", err)
	}

	return nil
}

func (s *CommentStore) DeleteComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM comments WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}

	return nil
}
