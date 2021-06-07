package domain

type AccountType string

const (
	AccountTypeUnknown      AccountType = ""
	AccountTypeUser         AccountType = "user"
	AccountTypeOrganization AccountType = "organization"
	AccountTypeEnterprise   AccountType = "enterprise"
	AccountTypeBot          AccountType = "bot"
)
