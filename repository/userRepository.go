package repository

import (
	"context"
	"github.com/disco07/blog/models"
	"time"
)

func (r Repository) CreateUser(user models.User, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `INSERT INTO user (full_name, username, email, password) VALUES ($1, $2, $3, $4)`

	stmt, err := r.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = r.DB.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	return nil
}
