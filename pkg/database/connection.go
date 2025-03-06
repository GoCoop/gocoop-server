package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {

	var (
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbName     = os.Getenv("DB_NAME")
	)

	encodedPw := url.QueryEscape(dbPassword)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, encodedPw, dbHost, dbPort, dbName)

	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	log.Println("> Connection successful!")

	return dbpool
}
