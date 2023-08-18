package common

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password string) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(pwdHash), nil
}

func CompareHashAndPassword(pwdHash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(password)) == nil
}
