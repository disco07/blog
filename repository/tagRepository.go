package repository

import (
	"context"
	"github.com/disco07/blog/models"
	"time"
)

func (r *Repository) FindTagById(id int, ctx context.Context) (*models.Tag, error) {
	var tag models.Tag
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT id, name FROM tag WHERE id = $1`
	err := r.DB.GetContext(ctx, &tag, query, id)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r Repository) FindAllTag(ctx context.Context) ([]*models.Tag, error) {
	var tags []*models.Tag
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT id, name FROM tag`
	err := r.DB.SelectContext(ctx, &tags, query)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (r Repository) InsertTag(tag models.Tag, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `INSERT INTO tag (name) VALUES ($1)`

	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = r.DB.NamedExecContext(ctx, query, tag)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) UpdateTag(tag models.Tag, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `UPDATE tag SET name = $1 WHERE id = $2`

	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = r.DB.NamedExecContext(ctx, query, tag)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) DeleteTag(id int, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `DELETE FROM tag WHERE id = $1`
	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = r.DB.NamedExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
