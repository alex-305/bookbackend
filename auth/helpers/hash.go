package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}

func ComparePassword(password string, dbpass string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbpass), []byte(password))
}
