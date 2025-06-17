package models

import "time"

type Feed struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Item struct {
	ID          string    `json:"id"`
	FeedID      string    `json:"feed_id"`
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Summary     string    `json:"summary"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Read        bool      `json:"read"`
}
