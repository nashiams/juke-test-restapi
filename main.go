package main

import (
	"juke-test-restapi/db"
	"juke-test-restapi/router"
	"log"
)

func main() {
    if err := db.Connect(); err != nil {
        log.Fatal(err)
    }
    defer db.DB.Close()

    r := router.SetupRouter()
    r.Run(":8080")
}
