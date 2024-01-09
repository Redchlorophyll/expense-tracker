package internal

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:expense-tracker1234@db.klwzjwiaupfalozgcpco.supabase.co:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
