package models

import "time"

type Post struct {
	ID          int       `json:"id" db:"id"`
	AuthorID    int       `json:"author_id" db:"author_id"`
	Title       string    `json:"title" db:"title"`
	Slug        string    `json:"slug" db:"slug"`
	Summary     string    `json:"summary" db:"summary"`
	Content     string    `json:"content" db:"content"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	Comments    Comment   `json:"-"`
}
