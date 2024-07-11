package domain

import "time"

type Item struct {
	ID          string     `json:"id,omitempty"`
	Code        string     `json:"code"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Stock       int        `json:"stock"`
	Available   bool       `json:"available"`
	Categories  []string   `json:"categories"`
	Provider    Provider   `json:"provider"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
