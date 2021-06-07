package bot

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/estafette/estafette-bot-github/clients/githubapi"
	"github.com/estafette/estafette-bot-github/domain"
	"github.com/rs/zerolog/log"
)

type Service interface {
	Run(ctx context.Context, githubEvent, githubDelivery, githubEventBody string) (err error)
}

func NewService(githubapiClient githubapi.Client) Service {
	return &service{
		githubapiClient: githubapiClient,
	}
}

type service struct {
	githubapiClient githubapi.Client
}

func (s *service) Run(ctx context.Context, githubEvent, githubDelivery, githubPayload string) (err error) {

	log.Info().Msgf("Running bot for event '%v'...", githubEvent)

	switch githubEvent {
	case "issues":

		var issuesEvent domain.IssuesEvent
		err = json.Unmarshal([]byte(githubPayload), &issuesEvent)
		if err != nil {
			return
		}

		switch issuesEvent.Action {
		case domain.IssuesEventActionOpened:
			if issuesEvent.Issue == nil {
				return fmt.Errorf("expected issue not to be nil for action '%v' for event '%v'", issuesEvent.Action, githubEvent)
			}

			// create welcome message in issue
			err = s.githubapiClient.AddCommentToIssue(ctx, *issuesEvent.Issue, domain.Comment{
				Body: "Thanks for raising an issue, we'll take a look at it soon",
			})
			if err != nil {
				return
			}

		default:
			return fmt.Errorf("action '%v' for event '%v' not supported", issuesEvent.Action, githubEvent)
		}

	case "commit_comment",
		"create",
		"delete",
		"deployment",
		"deployment_status",
		"fork",
		"gollum",
		"installation",
		"installation_repositories",
		"issue_comment",
		"label",
		"marketplace_purchase",
		"member",
		"membership",
		"milestone",
		"organization",
		"org_block",
		"page_build",
		"project_card",
		"project_column",
		"project",
		"public",
		"pull_request_review_comment",
		"pull_request_review",
		"pull_request",
		"push",
		"release",
		"repository",
		"status",
		"team",
		"team_add",
		"watch",
		"integration_installation_repositories":

	default:
		return fmt.Errorf("event '%v' not supported", githubEvent)
	}

	return nil
}
