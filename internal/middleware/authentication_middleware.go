package middleware

import (
	"github.com/EraldCaka/github-authentication-app/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorization(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/oauth" || ctx.Request.URL.Path == "/oauth/clientID" {
		ctx.Next()
		return
	}
	cookie, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		ctx.Abort()
		return
	}

	_, err = services.GetGitHubUser(cookie)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid authorization token!"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
