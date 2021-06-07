package githubapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/estafette/estafette-bot-github/clients/credentials"
	"github.com/estafette/estafette-bot-github/domain"
	"github.com/rs/zerolog/log"
	"github.com/sethgrid/pester"
)

// githubapi.Client allows to communicate with the Github api
type Client interface {
	AddCommentToIssue(ctx context.Context, issue domain.Issue, comment domain.Comment) (err error)
}

func NewClient(credentialsClient credentials.Client) Client {
	return &client{
		credentialsClient: credentialsClient,
	}
}

type client struct {
	credentialsClient credentials.Client
}

func (c *client) AddCommentToIssue(ctx context.Context, issue domain.Issue, comment domain.Comment) (err error) {

	// https://docs.github.com/en/rest/reference/issues#create-an-issue-comment
	_, err = c.callGithubAPI("POST", issue.CommentsURL, comment)
	if err != nil {
		return
	}

	return nil
}

func (c *client) callGithubAPI(method, url string, params interface{}) (body []byte, err error) {

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
	accessToken, err := c.credentialsClient.GetAccessToken()
	if err != nil {
		return
	}

	request.Header.Add("Authorization", fmt.Sprintf("token %v", accessToken))
	request.Header.Add("Accept", "application/vnd.github.v3+json")

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
