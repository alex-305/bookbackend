package helpers

import (
	"regexp"
)

func sanitize(s, rgx string) string {
	pattern := regexp.MustCompile(rgx)
	return pattern.ReplaceAllString(s, "")
}

func validate(s, rgx string) bool {
	regex, err := regexp.Compile(rgx)
	if err != nil {
		return false
	}

	return regex.MatchString(s)
}

func ValidateUsername(username string) bool {
	return validate(username, "^[a-zA-Z0-9_.]*$")
}
