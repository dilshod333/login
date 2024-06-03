package postgres

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	password = "dilshod"
	user     = "postgres"
	dbname   = "practise"
)

func Initialize() (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%d password=%s user=%s dbname=%s sslmode=disable", host, port, password, user, dbname)
	db, err := sql.Open("postgres", dbInfo)

	err = db.Ping()
	if err != nil {
		log.Fatal("Error while connecting...", err)
	}

	return db, nil
}


