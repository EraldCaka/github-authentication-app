package types

import (
	"log"
	"time"
)

type Commit struct {
	SHA    string `json:"sha"`
	NodeID string `json:"node_id"`
	Commit struct {
		Author struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"committer"`
		Message      string       `json:"message"`
		Tree         Tree         `json:"tree"`
		URL          string       `json:"url"`
		CommentCount int          `json:"comment_count"`
		Verification Verification `json:"verification"`
	} `json:"commit"`
	URL         string     `json:"url"`
	HTMLURL     string     `json:"html_url"`
	CommentsURL string     `json:"comments_url"`
	Author      UserCommit `json:"author"`
	Committer   UserCommit `json:"committer"`
	Parents     []Parent   `json:"parents"`
}

type Tree struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

type Verification struct {
	Verified  bool   `json:"verified"`
	Reason    string `json:"reason"`
	Signature string `json:"signature"`
	Payload   string `json:"payload"`
}

type UserCommit struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Parent struct {
	SHA     string `json:"sha"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

type CommitDB struct {
	ID        int       `json:"id,omitempty"`
	CommitSHA string    `json:"commit_sha"`
	Date      time.Time `json:"date"`
	RepoID    int       `json:"repo_id"`
}

type CommitReq struct {
	CommitSHA string    `json:"commit_sha"`
	Date      time.Time `json:"date"`
	RepoID    int       `json:"repo_id"`
}

type CommitRes struct {
	CommitSHA string    `json:"commit_sha"`
	Date      time.Time `json:"date"`
	RepoID    int       `json:"repo_id"`
}

func DateParsing(dateStr string) time.Time {
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		log.Printf("Error parsing date: %v\n", err)
		return time.Time{}
	}
	return date
}
