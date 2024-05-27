package main

import (
	"github.com/alex-305/bookbackend/internal/api"
	"github.com/alex-305/bookbackend/internal/db"
	_ "github.com/lib/pq"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	server := api.CreateServer("0.0.0.0:8080", db)

	err = server.Start()

	if err != nil {
		panic(err)
	}
}
