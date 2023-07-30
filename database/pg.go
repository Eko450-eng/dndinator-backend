package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func Init() {
	port := 5432
	user := os.Getenv("USER")
	host := os.Getenv("HOST")
	pass := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, dbname)

	DB, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	Db = DB
}
