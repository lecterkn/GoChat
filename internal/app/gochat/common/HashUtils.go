package common

import (
	"golang.org/x/crypto/bcrypt"
)

const HASH_COST = 14

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), HASH_COST)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func HashEquals(password string, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, []byte(password)) == nil
}
