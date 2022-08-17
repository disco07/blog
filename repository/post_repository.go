package repository

import (
	"context"
	"github.com/disco07/blog/models"
	"log"
	"time"
)

func (r *Repository) FindById(id int) (*models.Post, error) {
	var post models.Post
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT * posts WHERE id = $1`
	err := r.DB.GetContext(ctx, &post, query, id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &post, nil
}
