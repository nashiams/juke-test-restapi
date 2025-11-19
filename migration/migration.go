package main

import (
	"context"
	"fmt"
	"juke-test-restapi/db"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func createDatabaseIfNotExists() error {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using system env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/postgres?sslmode=disable",
		user, pass, host, port,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres database: %w", err)
	}
	defer pool.Close()

	var exists bool
	err = pool.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check database existence: %w", err)
	}

	if !exists {
		_, err = pool.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
		log.Printf("Database '%s' created successfully", dbName)
	} else {
		log.Printf("Database '%s' already exists", dbName)
	}

	return nil
}

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
	
	// Create database if it doesn't exist
	if err := createDatabaseIfNotExists(); err != nil {
		log.Fatal("Failed to create database:", err)
	}

	// Connect to the target database
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.DB.Close()

	RunMigration()
	
	log.Println("Migration completed successfully!")
}
