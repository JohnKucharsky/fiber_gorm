package db

import (
	"context"
	"github.com/induzo/gocom/database/pginit/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func New() *pgxpool.Pool {
	if err := godotenv.Load(); err != nil {
		log.Println("Can't load env")
	}

	dbAddress := "postgres://postgres:pass@db:5432/data?sslmode=disable"
	dbAddressEnv := os.Getenv("DB_URL")
	var dbAddressString = dbAddressEnv
	if dbAddressEnv == "" {
		dbAddressString = dbAddress
	}

	ctx := context.Background()
	pgi, err := pginit.New(dbAddressString)

	if err != nil {
		log.Fatal("Can't connect to db", err.Error())
	}

	pool, err := pgi.ConnPool(ctx)

	if err != nil {
		log.Fatal("Can't connect to db", err.Error())
	}

	return pool
}
