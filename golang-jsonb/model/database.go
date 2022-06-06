package model

import (
	"database/sql"
	"fmt"

	helper "golang-jsonb/helper"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "db_people"
)

func GetConnection() *sql.DB {
	// connection string
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	helper.CheckErrorPanic(err)

	return db
}
