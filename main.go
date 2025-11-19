package main

import (
	"juke-test-restapi/db"
	"juke-test-restapi/router"
	"log"

	"go.uber.org/zap"
)

func main() {
	// Initialize zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	r := router.SetupRouter()
	r.Run(":8080")
}
