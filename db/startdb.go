package db

import (
	"database/sql"
	"fmt"
	"os"
)

type DB struct {
	*sql.DB
}

func Start() (*DB, error) {
	dsn := "user=%s password=%s host=%s dbname=%s sslmode=disable"
	dsn = fmt.Sprintf(dsn, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) close() error {
	return db.DB.Close()
}
