package models

import "time"

type Todo struct {
	TodoId       string
	Task         string
	CreatedAt    time.Time
	IsCompleated bool
	CompleatedAt time.Time
	UserId       string
}
