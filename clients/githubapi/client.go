package githubapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/sethgrid/pester"
)

// githubapi.Client allows to communicate with the Github api
type Client interface {
}

func NewClient(accessToken string) Client {
	return &client{
		accessToken: accessToken,
	}
}

type client struct {
	accessToken string
}

func (c *client) callGithubAPI(method, url string, params interface{}, authorizationType, token string) (body []byte, err error) {

	// convert params to json if they're present
	var requestBody io.Reader
	if params != nil {
		data, err := json.Marshal(params)
		if err != nil {
			return body, err
		}
		requestBody = bytes.NewReader(data)
	}

	// create client, in order to add headers
	client := pester.New()
	client.MaxRetries = 3
	client.Backoff = pester.ExponentialJitterBackoff
	client.KeepLog = true
	request, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return
	}

	// add headers
	request.Header.Add("Authorization", fmt.Sprintf("%v %v", authorizationType, token))
	request.Header.Add("Accept", "application/vnd.github.machine-man-preview+json")

	// perform actual request
	response, err := client.Do(request)
	if err != nil {
		return
	}

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	// unmarshal json body
	var b interface{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Info().Err(err).Str("body", string(body)).Msgf("Deserializing response for '%v' Github api call failed", url)
		return
	}

	log.Info().Msgf("Received successful response for '%v' Github api call with status code %v", url, response.StatusCode)

	return
}
