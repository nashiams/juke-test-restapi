package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func Connect() error {
    // load .env
    if err := godotenv.Load(); err != nil {
        log.Println(".env not found, using system env")
    }

    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    name := os.Getenv("DB_NAME")

    // DSN
    dsn := fmt.Sprintf(
        "postgresql://%s:%s@%s:%s/%s",
        user, pass, host, port, name,
    )

    // create connection pool
    pool, err := pgxpool.New(context.Background(), dsn)
    if err != nil {
        log.Fatal("failed to create pool:", err)
    }

    // optional timeout for ping
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := pool.Ping(ctx); err != nil {
        return fmt.Errorf("failed to ping db: %w", err)
    }

    DB = pool
    log.Println("DB connected")
    return nil
}
