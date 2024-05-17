package booksvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Get(volumeid string, d *db.DB) (models.BookStats, error) {
	return d.GetBookStats(volumeid)
}
