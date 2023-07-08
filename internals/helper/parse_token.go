package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenWithClaims(tokenString, subject, username string, id int) (*jwt.Token, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invlaid token signing method")
		}
		return secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("token signature invalid: %v", err)
		}
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token not valid")
	}
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		return nil, errors.New("token already expired")
	}
	if claims.Subject != subject {
		return nil, errors.New("token invalid: subject mismatch")
	}
	if claims.Username != username {
		return nil, errors.New("token invalid: username mismatch")
	}
	if claims.ID != id {
		return nil, errors.New("token invalid: id mismatch")
	}
	return token, nil
}
