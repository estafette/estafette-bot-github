package domain

type Organization struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	Description string `json:"description"`
}
