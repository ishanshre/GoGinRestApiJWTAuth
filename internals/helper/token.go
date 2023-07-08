package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ishanshre/GoRestApiExample/internals/models"
)

var (
	AccessExpiresAt  = jwt.NewNumericDate(time.Now().Add(time.Minute * 15))
	RefreshExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour))
	IssuedAt         = jwt.NewNumericDate(time.Now())
	NotBefore        = jwt.NewNumericDate(time.Now())
	secret           = []byte(os.Getenv("jwt_secret"))
)

type Claims struct {
	Username string
	ID       int
	jwt.RegisteredClaims
}

func GenerateLoginResponse(id int, username string) (*models.LoginResponse, error) {
	access_token, err := generateAccessToken(id, username)
	if err != nil {
		return nil, err
	}
	refresh_token, err := generateRefreshToken(id, username)
	if err != nil {
		return nil, err
	}
	return &models.LoginResponse{
		Username:     username,
		ID:           id,
		AccessToken:  access_token,
		RefershToken: refresh_token,
	}, nil
}

func generateAccessToken(id int, username string) (string, error) {
	// creating access claim
	access_claims := &Claims{
		Username: username,
		ID:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: AccessExpiresAt,
			IssuedAt:  IssuedAt,
			NotBefore: NotBefore,
			Subject:   "access_token",
		},
	}

	// creating a new token with access claims
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)

	// sign the token with our unique secret from the environment files
	signedAccessToken, err := access_token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedAccessToken, nil

}

func generateRefreshToken(id int, username string) (string, error) {
	// creating custom refresh token claims
	refresh_claims := &Claims{
		Username: username,
		ID:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: RefreshExpiresAt,
			IssuedAt:  IssuedAt,
			NotBefore: NotBefore,
			Subject:   "refresh_token",
		},
	}

	// creating a new token with refresh claims
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)

	// sign the token with the secret
	signedRefreshToken, err := refresh_token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedRefreshToken, nil

}
