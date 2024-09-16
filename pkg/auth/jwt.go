// File: jwt.go
// Creation: Fri Aug 16 17:04:38 2024
// Time-stamp: <2024-08-29 09:44:59>
// Copyright (C): 2024 Pierre Lecocq

package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTConfig type struct
type JWTAuthConfig struct {
	Secret  string `mapstructure:"secret"`
	Issuer  string `mapstructure:"issuer"`
	Expires int    `mapstructure:"expires"`
	Disable bool   `mapstructure:"disable"`
	Debug   bool   `mapstructure:"debug"`
}

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// Generate a JWT token against a secret string that encapsulate a piece of identity (can be a username, an email, ...)
func GenerateJWTToken(userID int64, issuer string, secret string, expires int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		userID,
		jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := claims.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Validates a JWT token against the secret string used to generate it
func ValidateJWTToken(tokenString string, secret string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("Invalid token")
	}

	return nil
}

// Validate and decode a JWT token
func ValidateAndDecodeJWTToken(tokenString string, secret string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return 0, errors.New("Invalid claims")
	}

	return claims.UserID, nil
}
