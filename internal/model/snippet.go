package model

import (
	"time"
)

type Snippet struct {
	ID        int64         `json:"id"`
	Text      string        `json:"text"`
	Language  Language      `json:"language"`
	Title     string        `json:"title"`
	Author    User          `json:"author"`
	ShortURL  string        `json:"short_url"`
	CreatedAt time.Time     `json:"created_at"`
	TTL       time.Duration `json:"ttl"`
	Views     int64         `json:"views"`
	Comments  []Comment     `json:"comments"`
}
