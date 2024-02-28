package handlers

import (
	"github.com/EraldCaka/github-authentication-app/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserByID(ctx *gin.Context, dbConn *db.Postgres) {
	userID := ctx.Param("userID")
	user := dbConn.GetUserByID(ctx, userID)
	if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Failed to get user from DB"})

	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
}
func GetRepositories(ctx *gin.Context, dbConn *db.Postgres) {
	repositories, err := dbConn.GetRepositories(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Failed to get repositories from DB"})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "repositories": repositories})
}
func GetRepoByID(ctx *gin.Context, dbConn *db.Postgres) {
	userID := ctx.Param("repoID")
	repository := dbConn.GetRepoByID(ctx, userID)
	if repository == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Failed to get repo from DB"})

	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "repository": repository})
}
func GetCommits(ctx *gin.Context, dbConn *db.Postgres) {
	commits, err := dbConn.GetCommits(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Failed to get commits from DB"})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "repositories": commits})
}
func GetCommitByID(ctx *gin.Context, dbConn *db.Postgres) {
	commitID := ctx.Param("commitID")
	commit := dbConn.GetCommitByID(ctx, commitID)
	if commit == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Failed to get commit from DB"})

	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "repository": commit})
}
func GetCommitByRepoID(ctx *gin.Context, dbConn *db.Postgres) {
	repoID := ctx.Param("repoID")
	commits, err := dbConn.GetCommitsByRepoID(ctx, repoID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Failed to get commits from DB"})
	}
	ctx.JSON(http.StatusOK, commits)

}
