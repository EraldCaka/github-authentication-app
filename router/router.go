package router

import (
	"context"
	"github.com/EraldCaka/github-authentication-app/db"
	"github.com/EraldCaka/github-authentication-app/internal/handlers"
	"github.com/EraldCaka/github-authentication-app/internal/middleware"
	"github.com/EraldCaka/github-authentication-app/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()
	dbConn, err := db.NewPGInstance(context.Background())
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
		return
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Use(middleware.Authorization)
	r.GET("/oauth", handlers.GitHubOAuth)
	r.GET("/oauth/logout", handlers.LogoutUser)
	r.GET("/oauth/user", handlers.GithubUser)
	r.GET("/oauth/staredRepositories", handlers.UserStarredRepositories)
	r.GET("/oauth/repository/commit", handlers.RepositoryCommits)

	r.GET("/oauth/db/user/:userID", func(ctx *gin.Context) {
		handlers.GetUserByID(ctx, dbConn)
	})
	r.GET("/oauth/db/repo/:repoID", func(ctx *gin.Context) {
		handlers.GetRepoByID(ctx, dbConn)
	})
	r.GET("/oauth/db/commit/:commitID", func(ctx *gin.Context) {
		handlers.GetCommitByID(ctx, dbConn)
	})
	r.GET("/oauth/db/repositories", func(ctx *gin.Context) {
		handlers.GetRepositories(ctx, dbConn)
	})
	r.GET("/oauth/db/commits", func(ctx *gin.Context) {
		handlers.GetCommits(ctx, dbConn)
	})
	r.GET("/oauth/clientID", func(ctx *gin.Context) {
		if util.CLIENT_ID == "" {
			ctx.JSON(http.StatusUnauthorized, "Error finding clientID")
			return
		}
		ctx.JSON(http.StatusOK, util.CLIENT_ID)
	})
}

func Start(addr string) error {
	return r.Run(addr)
}
