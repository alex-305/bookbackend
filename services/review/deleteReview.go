package review

import "github.com/alex-305/bookbackend/db"

func DeleteReview(tok, reviewid string, db *db.DB) error {
	db.DeleteReview(reviewid)
	return nil
}
