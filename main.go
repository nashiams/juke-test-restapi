package main

import (
	"juke-test-restapi/db"
	_ "juke-test-restapi/docs"
	"juke-test-restapi/router"
	"log"

	"go.uber.org/zap"
)

// @title Employee Management API
// @version 1.0
// @description REST API for managing employee data
// @BasePath /api
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
