package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EraldCaka/github-authentication-app/util"
	"io"
	"net/http"
	"net/url"
	"time"
)

type AuthToken struct {
	Access_token string
}

type UserResult struct {
	Name  string
	Photo string
}

func GetGitHubOauthToken(code string) (*AuthToken, error) {
	const rootURl = "https://github.com/login/oauth/access_token"

	values := url.Values{}
	values.Add("code", code)
	values.Add("client_id", util.CLIENT_ID)
	values.Add("client_secret", util.CLIENT_SECRET)

	query := values.Encode()

	queryString := fmt.Sprintf("%s?%s", rootURl, bytes.NewBufferString(query))

	req, err := http.NewRequest("POST", queryString, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

	tokenBody := &AuthToken{
		Access_token: parsedQuery["access_token"][0],
	}

	return tokenBody, nil
}

func GetGitHubUser(accessToken string) (*UserResult, error) {
	rootUrl := "https://api.github.com/user"

	req, err := http.NewRequest("GET", rootUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)

	if err != nil {
		return nil, err
	}

	var GitHubUserRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GitHubUserRes); err != nil {
		return nil, err
	}

	userBody := &UserResult{
		Name:  GitHubUserRes["login"].(string),
		Photo: GitHubUserRes["avatar_url"].(string),
	}

	return userBody, nil
}
