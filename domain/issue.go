package domain

import "time"

type Issue struct {
	Labels    []Label    `json:"labels"`
	State     string     `json:"state"`
	Locked    bool       `json:"locked"`
	User      *Account   `json:"user"`
	Assignee  *Account   `json:"assignee"`
	Assignees []Account  `json:"assignees"`
	Milestone *Milestone `json:"milestone"`
	Comments  int        `json:"comments"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	ClosedAt  *time.Time `json:"closed_at"`

	CommentsURL string `json:"comments_url"`
}
