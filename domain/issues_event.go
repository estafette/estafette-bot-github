package domain

type IssuesEvent struct {
	Action       IssuesEventAction `json:"action"`
	Issue        *Issue            `json:"issue"`
	Milestone    *Milestone        `json:"milestone"`
	Repository   *Repository       `json:"repository"`
	Organization *Organization     `json:"organization"`
	Sender       *Account          `json:"sender"`
	Installation *Installation     `json:"installation"`
}
