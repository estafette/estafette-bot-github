package domain

type IssuesEventAction string

const (
	IssuesEventActionUnknown      IssuesEventAction = ""
	IssuesEventActionOpened       IssuesEventAction = "opened"
	IssuesEventActionEdited       IssuesEventAction = "edited"
	IssuesEventActionDeleted      IssuesEventAction = "deleted"
	IssuesEventActionTransferred  IssuesEventAction = "transferred"
	IssuesEventActionClosed       IssuesEventAction = "closed"
	IssuesEventActionReopened     IssuesEventAction = "reopened"
	IssuesEventActionLocked       IssuesEventAction = "locked"
	IssuesEventActionUnlocked     IssuesEventAction = "unlocked"
	IssuesEventActionPinned       IssuesEventAction = "pinned"
	IssuesEventActionUnpinned     IssuesEventAction = "unpinned"
	IssuesEventActionAssigned     IssuesEventAction = "assigned"
	IssuesEventActionUnassigned   IssuesEventAction = "unassigned"
	IssuesEventActionLabeled      IssuesEventAction = "labeled"
	IssuesEventActionUnlabeled    IssuesEventAction = "unlabeled"
	IssuesEventActionMilestoned   IssuesEventAction = "milestoned"
	IssuesEventActionDemilestoned IssuesEventAction = "demilestoned"
)
