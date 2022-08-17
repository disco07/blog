package models

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TagPost struct {
	PostID int `json:"post_id"`
	TagID  int `json:"tag_id"`
}
