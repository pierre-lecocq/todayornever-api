// File: jwt_test.go
// Creation: Fri Aug 16 17:05:32 2024
// Time-stamp: <2024-09-03 15:40:00>
// Copyright (C): 2024 Pierre Lecocq

package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWTToken(t *testing.T) {
	tokenString, _ := GenerateJWTToken(1, "testissuer", "secret", 1)

	// OK valid
	assert.Nil(t, ValidateJWTToken(tokenString, "secret"))

	// KO invalid
	assert.Error(t, ValidateJWTToken(tokenString, "badsecret"), "Invalid token")
	assert.Error(t, ValidateJWTToken("badtoken", "secret"), "Invalid token")
	assert.Error(t, ValidateJWTToken("badtoken", "badsecret"), "Invalid token")

	// OK valid
	res, err := ValidateAndDecodeJWTToken(tokenString, "secret")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), res)

	// KO invalid
	res, err = ValidateAndDecodeJWTToken("badtoken", "badsecret")
	assert.Error(t, err)
	assert.Equal(t, int64(0), res)
}
