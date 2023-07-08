package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func ComparePasswordHash(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
