package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("secret")

func ParseToken(r *http.Request) (string, error) {
	tokenStr := r.URL.Query().Get("token")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_id"].(string), nil
	}

	return "", err
}