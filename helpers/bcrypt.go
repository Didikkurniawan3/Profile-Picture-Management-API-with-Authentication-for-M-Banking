package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashResult, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashResult), nil
}

func ComparePassword(inputPass, dbPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(inputPass))
	return err == nil
}