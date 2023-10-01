package token

import (
	"fmt"
	"time"

	"github.com/edr3x/fiber-starter/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	Access  TokenType = "access"
	Refresh TokenType = "refresh"
)

func selectFunc(tokenType TokenType) (string, time.Duration) {
	var (
		secret    string
		expiresIn time.Duration
	)

	if tokenType == Access {
		secret = config.Env().Jwt.AccessSecret
		expiresIn = time.Minute * 15
		return secret, expiresIn
	}

	if tokenType == Refresh {
		secret = config.Env().Jwt.RefreshSecret
		expiresIn = time.Hour * 24 * 30
		return secret, expiresIn
	}

	return secret, expiresIn
}

func Generate(tokenType TokenType, userID string) (string, error) {
	secret, expiresIn := selectFunc(tokenType)

	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(expiresIn).Unix(),
	}
	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := userToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Verify(tokenString string, tokenType TokenType) (jwt.MapClaims, error) {
	secret, _ := selectFunc(tokenType)

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
