package handlers

import (
	"github.com/EraldCaka/github-authentication-app/internal/services"
	"github.com/EraldCaka/github-authentication-app/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "", "", false, false)
	util.ClearActiveToken()
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
func GitHubOAuth(ctx *gin.Context) {
	code := ctx.Query("code")

	if code == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}
	tokenRes, err := services.GetGitHubOauthToken(code)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": "Failed to retrieve OAuth token"})
		return
	}
	ctx.SetCookie("Authorization", tokenRes.Access_token, 60*60*24, "/", "localhost", false, false)
	util.SetActiveToken(tokenRes.Access_token)
	ctx.Redirect(http.StatusFound, util.ClientUserURI)
}
