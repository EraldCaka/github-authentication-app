package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EraldCaka/github-authentication-app/types"
	"github.com/EraldCaka/github-authentication-app/util"
	"io"
	"net/http"
	"time"
)

func GetGitHubUser(accessToken string) (*types.User, error) {

	req, err := http.NewRequest("GET", util.UserUrI, nil)

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

	var UserRes *types.User
	if err := json.NewDecoder(&resBody).Decode(&UserRes); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(UserRes)
	return UserRes, nil
}

func GetUserStarredRepos(accessToken string) ([]*types.Repository, error) {
	req, err := http.NewRequest("GET", util.StarredRepos, nil)
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
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var RepositoryRes []*types.Repository
	if err := json.NewDecoder(&resBody).Decode(&RepositoryRes); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(RepositoryRes)
	return RepositoryRes, nil
}
