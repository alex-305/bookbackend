package db

func (db DB) PostCommentLikes(user, commentid string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE comments SET likecount = likecount+1 WHERE commentid = $1`, commentid)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO user_likes_comment(username, commentid) VALUES($1,$2)`, user, commentid)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (db DB) DeleteCommentLikes(user, commentid string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE comments SET likecount = likecount-1 WHERE commentid = $1`, commentid)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`DELETE FROM user_likes_comment WHERE username = $1 AND commentid = $2`, user, commentid)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
