package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EraldCaka/github-authentication-app/db"
	"github.com/EraldCaka/github-authentication-app/internal/types"
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

	return UserRes, nil
}

func GetUserStarredRepos(accessToken string) ([]*types.Repository, error) {

	user, err := GetGitHubUser(accessToken)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", util.StarredRepos(user.Login), nil)
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
		return nil, errors.New("could not retrieve repositories")
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

	return RepositoryRes, nil
}

func GetRepositoryCommits(accessToken string, name string, repo string) ([]*types.Commit, error) {

	req, err := http.NewRequest("GET", util.RepoCommits(name, repo), nil)

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
		return nil, errors.New("could not retrieve commits")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var CommitResp []*types.Commit
	if err := json.NewDecoder(&resBody).Decode(&CommitResp); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return CommitResp, nil
}
func PopulateDBWithCurrentlyRegisteredUser(ctx context.Context, accessToken string, user *types.User, dbConn *db.Postgres) error {

	userReq := &types.UserReq{
		Name:     user.Login,
		Username: user.Name,
		ImgUrl:   user.AvatarURL,
	}
	userID, err := dbConn.CreateUser(ctx, userReq)

	if err != nil {
		fmt.Println(err, "database error user")
	}
	repos, err := GetUserStarredRepos(accessToken)

	if err != nil {
		return err
	}
	for _, repo := range repos {
		repoReq := &types.RepositoryReq{
			UserId:    userID,
			RepoName:  repo.Name,
			RepoOwner: repo.Owner.Login,
		}
		repoID, err := dbConn.CreateRepository(ctx, repoReq)

		if err != nil {
			fmt.Println(err, "database error repository")
		}
		commits, err := GetRepositoryCommits(accessToken, repoReq.RepoOwner, repoReq.RepoName)

		for _, commit := range commits {
			commitReq := &types.CommitReq{
				CommitSHA: commit.SHA,
				Date:      types.DateParsing(commit.Commit.Author.Date),
				RepoID:    repoID,
			}
			if err := dbConn.CreateCommit(ctx, commitReq); err != nil {
				fmt.Println(err, "database error commit")
			}
		}
	}
	return nil
}
