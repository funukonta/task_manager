package pkg

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func ConnectPostgres() (*sqlx.DB, error) {
	host := "task-manager"
	user := "postgres"
	password := "task-manager"
	dbname := "task-manager"
	ssl := "disable"

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, password, host, dbname, ssl)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Panicln("error connStr", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Panicln("error ping", err.Error())
	}

	return db, err
}
