package models

import "time"

type Post struct {
	ID          int
	Title       string
	Slug        string
	Summary     string
	Content     string
	PublishedAt time.Time
	Comments    Comment
}
