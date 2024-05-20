package db

import "errors"

func (db DB) PostReviewLikes(user, reviewid string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE reviews SET likecount = likecount+1 WHERE reviewid = $1`, reviewid)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO user_likes_review(username, reviewid) VALUES($1,$2)`, user, reviewid)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (db DB) DeleteReviewLikes(user, reviewid string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE reviews SET likecount = likecount-1 WHERE reviewid = $1`, reviewid)

	if err != nil {
		tx.Rollback()
		return err
	}

	result, err := tx.Exec(`DELETE FROM user_likes_review WHERE username = $1 AND reviewid = $2`, user, reviewid)

	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return errors.New("could not delete from user_likes_review")
	}

	tx.Commit()
	return err
}
