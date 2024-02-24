package router

import (
	"github.com/EraldCaka/github-authentication-app/internal/handlers"
	"github.com/EraldCaka/github-authentication-app/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5174" //react app
		},
		MaxAge: 12 * time.Hour,
	}))
	r.Use(middleware.Authorization)
	r.GET("/oauth", handlers.GitHubOAuth)
	r.GET("/oauth/logout", handlers.LogoutUser)
	r.GET("/oauth/user", handlers.GithubUser)
	r.GET("/oauth/staredRepositories", handlers.UserStarredRepositories)

}

func Start(addr string) error {
	return r.Run(addr)
}
