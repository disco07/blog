package repository

import (
	"context"
	"github.com/disco07/blog/models"
	"time"
)

func (r *Repository) FindPostById(id int) (*models.Post, error) {
	var post models.Post
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, author_id, title, slug, summary, content, published_at FROM post WHERE id = $1`
	err := r.DB.GetContext(ctx, &post, query, id)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r Repository) FindAllPost() ([]*models.Post, error) {
	var posts []*models.Post
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, author_id, title, slug, summary, content, published_at FROM post`
	err := r.DB.SelectContext(ctx, &posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r Repository) InsertMovie(post models.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO post (author_id, title, slug, summary, content, published_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.DB.NamedExecContext(ctx, query, post)
	if err != nil {
		return err
	}

	return nil
}