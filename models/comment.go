package models

import "time"

type Comment struct {
	ID          int
	PostID      int
	PublishedAt time.Time
}
