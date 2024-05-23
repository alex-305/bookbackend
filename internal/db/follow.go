package db

import "errors"

func (db *DB) PostFollow(follower, followed string) error {
	tx, err := db.Begin()

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO user_follows_user(follower, followed) VALUES($1, $2)`, follower, followed)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`UPDATE users SET followingcount = followingCount+1 WHERE username=$1`, follower)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`UPDATE users SET followercount = followercount+1 WHERE username=$1`, followed)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (db *DB) DeleteFollow(follower, followed string) error {
	tx, err := db.Begin()

	if err != nil {
		tx.Rollback()
		return err
	}
	result, err := tx.Exec(`DELETE FROM user_follows_user WHERE follower=$1 AND followed=$2`, follower, followed)

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
		return errors.New("could not unfollow")
	}

	_, err = tx.Exec(`UPDATE users SET followercount=followercount-1 WHERE username=$1`, followed)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`UPDATE users SET followingcount=followingcount-1 WHERE username=$1`, follower)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
