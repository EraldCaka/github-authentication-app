package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EraldCaka/github-authentication-app/db"
	"github.com/EraldCaka/github-authentication-app/internal/services"
	"github.com/EraldCaka/github-authentication-app/internal/types"
	"github.com/EraldCaka/github-authentication-app/util"
	"log"
	"net/http"
	"time"
)

func main() {
	util.InitEnvironmentVariables()
	dbConn, err := db.NewPGInstance(context.Background())
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
		return
	}
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	fmt.Println("worker has started...")
	runTask(dbConn)
	for range ticker.C {
		runTask(dbConn)
	}
}

func runTask(dbConn *db.Postgres) {
	token := util.ACTIVE_TOKEN
	if token == "" {
		fmt.Println("you are not logged in yet")
		return
	}
	commits, err := dbConn.GetCommits(context.Background())
	if err != nil {
		fmt.Println("There is an issue with getting commits")
		return
	}

	if len(commits) == 0 {
		githubUser, err := services.GetGitHubUser(token)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := services.PopulateDBWithCurrentlyRegisteredUser(context.Background(), token, githubUser, dbConn); err != nil {
			fmt.Println("There was an error on populating the tables: ", err)
			return
		}
		fmt.Println("Tables were successfully populated  with ", githubUser.Login, "'s data")
	} else {
		repos, err := dbConn.GetRepositories(context.Background())
		if err != nil {
			fmt.Println("couldn't get repository", err)
			return
		}

		for _, repo := range repos {
			lastDate, err := dbConn.GetLatestCommit(context.Background(), repo.ID)
			if err != nil {
				fmt.Println(err)
				continue
			}
			formattedDate := lastDate.Format("2006-01-02T15:04:05Z07:00")
			commitsURL := util.CommitsSince(repo.RepoOwner, repo.RepoName, formattedDate)
			resp, err := http.Get(commitsURL)
			if err != nil {
				fmt.Println("Error making HTTP request:", err)
				continue
			}
			defer resp.Body.Close()
			var commitsData []types.Commit
			err = json.NewDecoder(resp.Body).Decode(&commitsData)
			if err != nil {
				fmt.Println("Error decoding JSON response:", err)
				continue
			}

			for _, commitData := range commitsData {
				commitReq := &types.CommitReq{
					CommitSHA: commitData.SHA,
					Date:      types.DateParsing(commitData.Commit.Author.Date),
					RepoID:    repo.ID,
				}
				if err := dbConn.CreateCommit(context.Background(), commitReq); err != nil {
					fmt.Println("error committing commit with id ", commitData.SHA, "\n", err)
				}
			}
		}
		fmt.Println("Commits were updated successfully!")
	}
}
