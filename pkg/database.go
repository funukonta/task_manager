package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sqlx.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, password, host, dbname, ssl)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Panicln("error connStr", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Panicln("error ping", err.Error())
	}

	m, err := migrate.New("file://../migrate", connStr)
	if err != nil {
		log.Panicln("error migrate", err.Error())
	}
	// err = m.Force(2)
	// if err != nil {
	// 	log.Panicln("error migrate force", err.Error())
	// }

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Panicln("error migrate up", err.Error())
	}

	return db, nil
}
