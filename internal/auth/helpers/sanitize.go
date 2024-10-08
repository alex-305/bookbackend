package helpers

import (
	"regexp"

	"github.com/alex-305/bookbackend/internal/models"
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

func ValidateUsername(username models.Username) bool {
	return validate(string(username), "^[a-zA-Z0-9_.]*$")
}

func ValidateEmail(email string) bool {
	regex := "(?:[a-z0-9!#$%&'*+=?^_`{|}~-]+(?:.[a-z0-9!#$%&'*+=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9].{3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)]"
	return validate(email, regex)
}
