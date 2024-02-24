package handlers

import (
	"fmt"
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
	fmt.Println(repos)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "repos": repos})
}
