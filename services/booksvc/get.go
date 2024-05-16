package booksvc

import (
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func Get(volumeid string, d *db.DB) (models.BookStats, error) {
	return d.GetBookStats(volumeid)
}
