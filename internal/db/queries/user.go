package queries

func GetUser() string {
	return `SELECT u.username, u.email, u.description, u.join_date, u.followercount, u.followingcount, CASE WHEN ufu.follower IS NOT NULL THEN TRUE ELSE FALSE END AS isfollowing FROM users u LEFT JOIN user_follows_user ufu ON u.username=ufu.followed AND ufu.follower=$1 WHERE username = $2`
}
