package models

import "time"

type Comment struct {
	ID          int       `json:"id"`
	PostID      int       `json:"post_id"`
	AuthorID    int       `json:"author_id"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
}
