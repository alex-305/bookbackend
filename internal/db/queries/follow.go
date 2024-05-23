package queries

func GetFollower() string {
	return `SELECT uf1.follower, CASE WHEN uf2.follower IS NOT NULL THEN TRUE ELSE FALSE END AS followed FROM user_follows_user AS uf1 LEFT JOIN user_follows_user AS uf2 ON uf1.follower=uf2.followed AND uf2.follower=$1 WHERE uf1.followed=$2`
}

func GetFollowing() string {
	return `SELECT uf1.followed, CASE WHEN uf2.followed IS NOT NULL THEN TRUE ELSE FALSE END AS followed FROM user_follows_user AS uf1 LEFT JOIN user_follows_user uf2 ON uf1.followed=uf2.follower AND uf2.followed=$1 WHERE uf1.follower=$2`
}
