package domain

import "time"

type Comment struct {
	ID        int        `json:"id,omitempty"`
	User      Account    `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Body      string     `json:"body"`
}
