package domain

import "time"

type Milestone struct {
	ID           int        `json:"id"`
	Number       int        `json:"number"`
	Title        string     `json:"title"`
	Creator      Account    `json:"creator"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	State        string     `json:"state"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DueOn        *time.Time `json:"due_on"`
	ClosedAt     *time.Time `json:"closed_at"`
}
