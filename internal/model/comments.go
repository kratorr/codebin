package model

import "time"

type Comment struct {
	ID        int64     `json:"id"`
	Author    User      `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
