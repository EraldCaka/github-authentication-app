package services

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/EraldCaka/github-authentication-app/internal/types"
	"github.com/EraldCaka/github-authentication-app/util"
	"io"
	"net/http"
	"net/url"
	"time"
)

func GetGitHubOauthToken(code string) (*types.AuthToken, error) {
	values := url.Values{}
	values.Add("code", code)
	values.Add("client_id", util.CLIENT_ID)
	values.Add("client_secret", util.CLIENT_SECRET)

	query := values.Encode()

	queryString := fmt.Sprintf("%s?%s", util.TokenUrl, bytes.NewBufferString(query))

	req, err := http.NewRequest("POST", queryString, nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	var resBody bytes.Buffer

	_, err = io.Copy(&resBody, res.Body)

	if err != nil {
		return nil, err
	}

	parsedQuery, err := url.ParseQuery(resBody.String())

	if err != nil {
		return nil, err
	}

	tokenBody := &types.AuthToken{
		Access_token: parsedQuery["access_token"][0],
	}

	return tokenBody, nil
}
