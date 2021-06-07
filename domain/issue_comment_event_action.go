package domain

type IssueCommentEventAction string

const (
	IssueCommentEventActionUnknown IssueCommentEventAction = ""
	IssueCommentEventActionCreated IssueCommentEventAction = "created"
	IssueCommentEventActionEdited  IssueCommentEventAction = "edited"
	IssueCommentEventActionDeleted IssueCommentEventAction = "deleted"
)
