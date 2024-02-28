package db

import (
	"context"
	"fmt"
	"github.com/EraldCaka/github-authentication-app/internal/types"
	"log"
	"time"
)

func (pg *Postgres) CreateUser(ctx context.Context, u *types.UserReq) (int, error) {
	query := fmt.Sprintf("INSERT INTO public.users (name, username, img_url) VALUES ('%v','%v','%v') RETURNING id", u.Name, u.Username, u.ImgUrl)
	var userID int
	err := pg.db.QueryRow(ctx, query).Scan(&userID)
	if err != nil {
		log.Printf("Unable to insert user: %v\n", err)
		return 0, err
	}
	return userID, nil
}

func (pg *Postgres) CreateCommit(ctx context.Context, commit *types.CommitReq) error {
	formattedDateStr := commit.Date.Format("2006-01-02T15:04:05Z")
	query := fmt.Sprintf("INSERT INTO public.commits (commit_sha, date, repo_id) VALUES ('%v','%v','%v')RETURNING id", commit.CommitSHA, formattedDateStr, commit.RepoID)
	var commitID int
	err := pg.db.QueryRow(ctx, query).Scan(&commitID)
	if err != nil {
		log.Printf("Unable to insert commit: %v\n", err)
		return err
	}
	return nil
}

func (pg *Postgres) CreateRepository(ctx context.Context, repository *types.RepositoryReq) (int, error) {
	query := fmt.Sprintf("INSERT INTO public.starred_repos (user_id, repo_name, repo_owner) VALUES ('%v','%v','%v') RETURNING id", repository.UserId, repository.RepoName, repository.RepoOwner)

	var repoID int
	err := pg.db.QueryRow(ctx, query).Scan(&repoID)
	if err != nil {
		log.Printf("Unable to insert repository: %v\n", err)
		return 0, err
	}

	return repoID, nil
}

func (pg *Postgres) GetLatestCommit(ctx context.Context, repoID int) (*time.Time, error) {
	var latestDate time.Time

	query := fmt.Sprintf("SELECT MAX(date) FROM commits WHERE repo_id = %d", repoID)

	err := pg.db.QueryRow(ctx, query).Scan(&latestDate)
	if err != nil {
		log.Printf("Error retrieving latest commit date: %v\n", err)
		return nil, err
	}

	return &latestDate, nil
}

func (pg *Postgres) GetUserByID(ctx context.Context, userID string) *types.UserDB {
	query := fmt.Sprintf("SELECT * FROM public.users WHERE id = '%v'", userID)
	row := pg.db.QueryRow(ctx, query)
	var userDB types.UserDB

	err := row.Scan(&userDB.ID, &userDB.Name, &userDB.Username, &userDB.ImgUrl)
	if err != nil {
		log.Printf("Error scanning user data: %v\n", err)
		return &types.UserDB{}
	}

	return &userDB

}
func (pg *Postgres) GetRepositories(ctx context.Context) ([]types.RepositoryDB, error) {
	var repositories []types.RepositoryDB

	query := "SELECT * FROM public.starred_repos"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying repositories: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var repo types.RepositoryDB
		err := rows.Scan(&repo.ID, &repo.UserId, &repo.RepoName, &repo.RepoOwner)
		if err != nil {
			log.Printf("Error scanning repository row: %v\n", err)
			continue
		}
		repositories = append(repositories, repo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over repository rows: %v\n", err)
		return nil, err
	}

	return repositories, nil
}

func (pg *Postgres) GetRepoByID(ctx context.Context, repoID string) *types.RepositoryDB {
	query := fmt.Sprintf("SELECT * FROM public.starred_repos WHERE id = '%v'", repoID)
	row := pg.db.QueryRow(ctx, query)
	var repoDB types.RepositoryDB

	err := row.Scan(&repoDB.ID, &repoDB.RepoName, &repoDB.RepoOwner)
	if err != nil {
		log.Printf("Error scanning commit data: %v\n", err)
		return &types.RepositoryDB{}
	}
	return &repoDB
}

func (pg *Postgres) GetCommits(ctx context.Context) ([]types.CommitDB, error) {
	var commits []types.CommitDB

	query := "SELECT * FROM public.commits"
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying commits: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var commit types.CommitDB
		err := rows.Scan(&commit.ID, &commit.CommitSHA, &commit.Date, &commit.RepoID)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		commits = append(commits, commit)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over commit rows: %v\n", err)
		return nil, err
	}

	return commits, nil
}

func (pg *Postgres) GetCommitByID(ctx context.Context, commitID string) *types.CommitDB {
	query := fmt.Sprintf("SELECT * FROM public.commits WHERE id = '%v'", commitID)
	row := pg.db.QueryRow(ctx, query)
	var commit types.CommitDB

	err := row.Scan(&commit.ID, &commit.CommitSHA, &commit.Date, &commit.RepoID)
	if err != nil {
		log.Printf("Error scanning commit data: %v\n", err)
		return &types.CommitDB{}
	}
	return &commit
}
func (pg *Postgres) GetCommitsByRepoID(ctx context.Context, repoID string) ([]types.CommitDB, error) {
	var commits []types.CommitDB

	query := fmt.Sprintf("SELECT * FROM public.commits WHERE repo_id = %v", repoID)
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error querying commits: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var commit types.CommitDB
		err := rows.Scan(&commit.ID, &commit.CommitSHA, &commit.Date, &commit.RepoID)
		if err != nil {
			log.Printf("Error scanning commit row: %v\n", err)
			continue
		}
		commits = append(commits, commit)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over commit rows: %v\n", err)
		return nil, err
	}

	return commits, nil
}
