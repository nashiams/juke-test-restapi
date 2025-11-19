package main

import (
	"context"
	"juke-test-restapi/db"
	"log"
)

func RunMigration() {
    migrations := []string{
        `
        CREATE TABLE IF NOT EXISTS employees (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            position VARCHAR(100) NOT NULL,
            salary NUMERIC(12,2) NOT NULL,
            created_at TIMESTAMP DEFAULT NOW()
        );
        `,
    }

    for _, m := range migrations {
        _, err := db.DB.Exec(context.Background(), m)
        if err != nil {
            log.Fatal("migration error:", err)
        }
    }

    log.Println("migration done")
}

func main() {
	log.Println("Starting migration...")
	
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.DB.Close()

	RunMigration()
	
	log.Println("Migration completed successfully!")
}
