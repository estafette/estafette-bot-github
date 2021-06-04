package bot

import (
	"fmt"

	"github.com/estafette/estafette-bot-github/clients/githubapi"
	"github.com/rs/zerolog/log"
)

type Service interface {
	Run(githubEvent, githubEventBody string) (err error)
}

func NewService(githubapiClient githubapi.Client) Service {
	return &service{
		githubapiClient: githubapiClient,
	}
}

type service struct {
	githubapiClient githubapi.Client
}

func (s *service) Run(githubEvent, githubEventBody string) (err error) {

	log.Info().Msgf("Running bot for event '%v'...", githubEvent)

	switch githubEvent {
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
		"issues",
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
		return fmt.Errorf("Event '%v' not supported", githubEvent)
	}

	return nil
}
