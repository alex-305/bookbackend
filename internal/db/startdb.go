package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func Connect() (*DB, error) {
	dsn := "user=%s password=%s host=%s dbname=%s sslmode=disable"
	dsn = fmt.Sprintf(dsn, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("POSTGRES_DB"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Close() error {
	instance, err := db.DB.DB()
	if err != nil {
		return err
	}
	err = instance.Close()

	if err != nil {
		return err
	}
	return nil
}
