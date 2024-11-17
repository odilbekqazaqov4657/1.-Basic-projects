package models

import "time"

type Post struct {
	PostId    string
	Title     string
	Content   string
	CreatedAt time.Time
}
