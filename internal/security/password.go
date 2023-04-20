package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func ValidatePassword(hash *string, password string) bool {
	bHash := []byte(*hash)
	bPassword := []byte(password)
	return bcrypt.CompareHashAndPassword(bHash, bPassword) == nil
}