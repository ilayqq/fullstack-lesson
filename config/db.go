package config

import (
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	DB, err := sqlx.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	log.Println("Success connect Postgres")
}
