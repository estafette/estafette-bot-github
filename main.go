package main

import (
	"runtime"

	"github.com/alecthomas/kingpin"
	"github.com/estafette/estafette-bot-github/clients/credentials"
	"github.com/estafette/estafette-bot-github/clients/githubapi"
	"github.com/estafette/estafette-bot-github/services/bot"
	foundation "github.com/estafette/estafette-foundation"
	"github.com/rs/zerolog/log"
)

var (
	appgroup  string
	app       string
	version   string
	branch    string
	revision  string
	buildDate string
	goVersion = runtime.Version()
)

var (
	githubEvent     = kingpin.Flag("github-event", "Event that triggered this bot.").Envar("ESTAFETTE_TRIGGER_GITHUB_EVENT").String()
	githubEventBody = kingpin.Flag("github-event-body", "Event body that triggered this bot.").Envar("ESTAFETTE_TRIGGER_GITHUB_EVENT_BODY").String()
	credentialsPath = kingpin.Flag("credentials-path", "Path to file with Github api token credentials configured at the CI server, passed in to this trusted extension.").Default("/credentials/github_api_token.json").String()
)

func main() {

	// parse command line parameters
	kingpin.Parse()

	// init log format from envvar ESTAFETTE_LOG_FORMAT
	foundation.InitLoggingFromEnv(appgroup, app, version, branch, revision, buildDate)

	credentialsClient := credentials.NewClient(*credentialsPath)
	githubapiClient := githubapi.NewClient(credentialsClient)
	botService := bot.NewService(githubapiClient)

	err := botService.Run(githubEvent, githubEventBody)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed running bot")
	}

	log.Info().Msg("Finished running bot...")
}
