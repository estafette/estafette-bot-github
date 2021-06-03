package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"

	foundation "github.com/estafette/estafette-foundation"
	"github.com/rs/zerolog/log"
)

type Client interface {
	GetAccessToken(credentialsPath string) (accessToken string, err error)
}

func NewClient() Client {
	return &client{}
}

type client struct {
}

func (c *client) GetAccessToken(credentialsPath string) (accessToken string, err error) {

	// get api token from injected credentials
	var credentials []APITokenCredentials
	// use mounted credential file if present instead of relying on an envvar
	if runtime.GOOS == "windows" {
		credentialsPath = "C:" + credentialsPath
	}
	if foundation.FileExists(credentialsPath) {
		log.Info().Msgf("Reading credentials from file at path %v...", credentialsPath)
		credentialsFileContent, err := ioutil.ReadFile(credentialsPath)
		if err != nil {
			return "", err
		}
		err = json.Unmarshal(credentialsFileContent, &credentials)
		if err != nil {
			return "", err
		}
		if len(credentials) == 0 {
			return "", fmt.Errorf("Found 0 credentials in file %v", credentialsPath)
		}
		log.Debug().Msgf("Read %v credentials", len(credentials))
	}
	if len(credentials) == 0 {
		return "", fmt.Errorf("Found 0 credentials in file %v", credentialsPath)
	}

	return credentials[0].AdditionalProperties.Token, nil
}
