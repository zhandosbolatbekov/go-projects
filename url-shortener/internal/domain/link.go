package domain

import "time"

type Link struct {
	ID        int64     `json:"id"`
	URL       string    `json:"url"`
	ShortCode string    `json:"short_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewLink(ID int64, URL string, shortCode string, createdAt time.Time, updatedAt time.Time) *Link {
	return &Link{ID: ID, URL: URL, ShortCode: shortCode, CreatedAt: createdAt, UpdatedAt: updatedAt}
}
