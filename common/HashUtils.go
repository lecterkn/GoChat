package common

import "golang.org/x/crypto/bcrypt"

const HASH_COST = 14

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), HASH_COST)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func HashEquals(password string, hash []byte) bool {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, hash)
	return err != nil
}