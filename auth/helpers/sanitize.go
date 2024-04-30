package helpers

import (
	"errors"
	"regexp"
)

func sanitize(s, rgx string) string {
	pattern := regexp.MustCompile(rgx)
	return pattern.ReplaceAllString(s, "")
}

func ValidateUsername(username string) error {
	if sanitize(username, "^[a-zA-Z0-9_.]*$") != username {
		return errors.New("invalid username")
	}
	return nil
}
