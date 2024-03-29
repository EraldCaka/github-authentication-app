package main

import (
	"github.com/EraldCaka/github-authentication-app/router"
	"github.com/EraldCaka/github-authentication-app/util"
)

func main() {
	util.InitEnvironmentVariables()
	router.InitRouter()
	router.Start("0.0.0.0:5555")
}
