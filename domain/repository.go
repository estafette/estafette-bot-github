package domain

import "time"

type Repository struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	FullName        string     `json:"full_name"`
	Private         bool       `json:"private"`
	Owner           Account    `json:"owner"`
	Fork            bool       `json:"fork"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	PushedAt        *time.Time `json:"pushed_at"`
	StargazersCount int        `json:"stargazers_count"`
	WatchersCount   int        `json:"watchers_count"`
	HasIssues       bool       `json:"has_issues"`
	HasProjects     bool       `json:"hahas_projectss_issues"`
	HasDownloads    bool       `json:"has_downloads"`
	HasWiki         bool       `json:"has_wiki"`
	HasPages        bool       `json:"has_pages"`
	ForksCount      int        `json:"forks_count"`
	Archived        bool       `json:"archived"`
	Disabled        bool       `json:"disabled"`
	OpenIssuesCount int        `json:"open_issues_count"`
	DefaultBranch   string     `json:"default_branch"`
}
