package models

import "time"

type Comment struct {
	ID          int       `json:"id"`
	PostID      int       `json:"post_id"`
	PublishedAt time.Time `json:"published_at"`
}
