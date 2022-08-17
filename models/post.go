package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	Comments    Comment   `json:"-"`
}
