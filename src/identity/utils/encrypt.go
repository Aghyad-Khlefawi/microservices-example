package utils

import "golang.org/x/crypto/bcrypt"

func HashPasswword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err != nil {
		return "", nil

	}

	return string(hash), nil
}
