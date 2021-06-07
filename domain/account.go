package domain

type Account struct {
	Login     string      `json:"login"`
	ID        int         `json:"id"`
	Type      AccountType `json:"type"`
	SiteAdmin bool        `json:"site_admin"`
}
