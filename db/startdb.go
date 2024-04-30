package db

import (
	"database/sql"
	"net/url"
)

type DB struct {
	*sql.DB
}

func StartDB() (*DB, error) {
	dsn := url.URL{
		User:   url.UserPassword("blank", "blank"),
		Host:   "localhost:5432",
		Scheme: "postgres",
		Path:   "databasename",
	}

	db, err := sql.Open("postgres", dsn.String())

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
