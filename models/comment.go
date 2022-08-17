package models

import "time"

type Comment struct {
	ID          int       `json:"id" db:"id"`
	PostID      int       `json:"post_id" db:"post_id"`
	AuthorID    int       `json:"author_id" db:"author_id"`
	Content     string    `json:"content" db:"content"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
}
