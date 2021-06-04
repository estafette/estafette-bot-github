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
	GetAccessToken() (accessToken string, err error)
}

func NewClient(credentialsPath string) Client {
	if runtime.GOOS == "windows" {
		credentialsPath = "C:" + credentialsPath
	}

	return &client{
		credentialsPath: credentialsPath,
	}
}

type client struct {
	credentialsPath string
}

func (c *client) GetAccessToken() (accessToken string, err error) {

	// get api token from injected credentials
	var credentials []APITokenCredentials
	// use mounted credential file if present instead of relying on an envvar
	if foundation.FileExists(c.credentialsPath) {
		log.Info().Msgf("Reading credentials from file at path %v...", c.credentialsPath)
		credentialsFileContent, err := ioutil.ReadFile(c.credentialsPath)
		if err != nil {
			return "", err
		}
		err = json.Unmarshal(credentialsFileContent, &credentials)
		if err != nil {
			return "", err
		}
		if len(credentials) == 0 {
			return "", fmt.Errorf("Found 0 credentials in file %v", c.credentialsPath)
		}
		log.Debug().Msgf("Read %v credentials", len(credentials))
	}
	if len(credentials) == 0 {
		return "", fmt.Errorf("Found 0 credentials in file %v", c.credentialsPath)
	}

	return credentials[0].AdditionalProperties.Token, nil
}
