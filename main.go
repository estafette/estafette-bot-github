package main

import (
	"context"
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
	githubDelivery  = kingpin.Flag("github-event-delivery", "Event devliery id that triggered this bot.").Envar("ESTAFETTE_TRIGGER_GITHUB_DELIVERY").String()
	githubPayload   = kingpin.Flag("github-event-payload", "Event payload that triggered this bot.").Envar("ESTAFETTE_TRIGGER_GITHUB_PAYLOAD").String()
	credentialsPath = kingpin.Flag("credentials-path", "Path to file with Github api token credentials configured at the CI server, passed in to this trusted extension.").Default("/credentials/github_api_token.json").String()
)

func main() {

	// parse command line parameters
	kingpin.Parse()

	// init log format from envvar ESTAFETTE_LOG_FORMAT
	foundation.InitLoggingFromEnv(appgroup, app, version, branch, revision, buildDate)

	// create context to cancel commands on sigterm
	ctx := foundation.InitCancellationContext(context.Background())

	credentialsClient := credentials.NewClient(*credentialsPath)
	githubapiClient := githubapi.NewClient(credentialsClient)
	botService := bot.NewService(githubapiClient)

	err := botService.Run(ctx, *githubEvent, *githubDelivery, *githubPayload)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed running bot")
	}

	log.Info().Msg("Finished running bot...")
}
