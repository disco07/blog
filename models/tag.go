package models

type Tag struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type TagPost struct {
	PostID int `json:"post_id" db:"post_id"`
	TagID  int `json:"tag_id" db:"tag_id"`
}
