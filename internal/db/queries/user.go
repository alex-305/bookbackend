package queries

func FromUser() string {
	return `users u`
}

func SelectUserID() string {
	return `u.userid`
}

func SelectUsername() string {
	return `u.username`
}

func SelectUser() string {
	return `u.username, u.email, u.description, u.join_date, u.followercount, u.followingcount, CASE WHEN ufu.followerid IS NOT NULL THEN TRUE ELSE FALSE END AS isfollowing`
}

func SelectCredentials() string {
	return `u.username, u.email, u.password`
}

func JoinUserFollows() string {
	return ` LEFT JOIN user_follows_user ufu ON u.username=ufu.followedid AND ufu.followerid = ? `
}
