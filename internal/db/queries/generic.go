package queries

func JoinForUsername(tablename string) string {
	return ` LEFT JOIN users u ON ` + tablename + `.userid = u.userid `
}

func JoinForUserID(tablename string) string {
	return ` LEFT JOIN users u ON ` + tablename + `.username = ? `
}
