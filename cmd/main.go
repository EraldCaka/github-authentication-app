package main

import (
	"github.com/EraldCaka/github-authentication-app/router"
	"github.com/EraldCaka/github-authentication-app/util"
)

func main() {
	util.InitEnvironmentVariables()
	//dbConn, err := db.NewPGInstance(context.Background())
	//if err != nil {
	//	log.Fatalf("could not initialize database connection: %s", err)
	//}
	//fmt.Println(dbConn)

	router.InitRouter()
	router.Start("0.0.0.0:5000")
}
