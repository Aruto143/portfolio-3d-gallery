package db

import (
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB() *sqlx.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://appuser:apppass@localhost:5432/portfolio?sslmode=disable"
	}

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open postgres connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}

	return db
}