package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// GetDbConn tries to establish a connection to postgres and return the connection handler
func GetDbConn() (*sql.DB, error) {

	databaseURL := getTCP()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db, nil
}

func getTCP() string {
	return fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s sslmode=disable", "127.0.0.1", "postgres", "passw0rd!", "5432", "postgres")
}
