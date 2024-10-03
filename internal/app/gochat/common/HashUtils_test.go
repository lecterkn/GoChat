package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "mypassword"
	_, err := HashPassword(password)
	assert.NoError(t, err)
}

func TestHashEquals(t *testing.T) {
	password := "mypassword"
	hash, err := HashPassword(password)
	assert.Nil(t, err)
	assert.True(t, HashEquals(password, hash))
}
