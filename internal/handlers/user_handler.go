package handlers

import (
	"github.com/EraldCaka/github-authentication-app/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GithubUser(ctx *gin.Context) {
	cookie, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}
	user, err := services.GetGitHubUser(cookie)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Failed to get GitHub user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
}

func UserStarredRepositories(ctx *gin.Context) {
	cookie, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}
	repos, err := services.GetUserStarredRepos(cookie)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Failed to get user repositories!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "repos": repos})
}

func RepositoryCommits(ctx *gin.Context) {
	var name = ctx.Query("name")
	var repo = ctx.Query("repo")
	cookie, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}
	commits, err := services.GetRepositoryCommits(cookie, name, repo)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Failed to get repositories commits!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "commits": commits})
}
