package util

import (
	"fmt"
)

// const RedirectLink = "http://localhost:5000/oauth"
// const AuthUrl = "https://github.com/login/oauth/authorize"
const TokenUrl = "https://github.com/login/oauth/access_token"
const UserUrI = "https://api.github.com/user"
const ClientUserURI = "http://localhost:5173/user"

var CommitsSince = func(name string, repo string, since string) string {
	return fmt.Sprintf("https://api.github.com/repos/%v/%v/commits?since=%v", name, repo, since)
}

var StarredRepos = func(name string) string {
	return fmt.Sprintf("https://api.github.com/users/%v/starred", name)
}
var RepoCommits = func(name string, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%v/%v/commits", name, repo)
}
