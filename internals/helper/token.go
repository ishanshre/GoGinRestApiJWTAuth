package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ishanshre/GoRestApiExample/internals/models"
	uuid "github.com/satori/go.uuid"
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

type Token struct {
	AccessToken  *TokenDetail
	RefreshToken *TokenDetail
}

func GenerateLoginResponse(id int, username string) (*models.LoginResponse, *Token, error) {
	tokenID := uuid.NewV4().String()
	access_token_detail, err := generateAccessToken(id, username, tokenID)
	if err != nil {
		return nil, nil, err
	}
	refresh_token_detail, err := generateRefreshToken(id, username, tokenID)
	if err != nil {
		return nil, nil, err
	}
	return &models.LoginResponse{
			Username:     username,
			ID:           id,
			AccessToken:  *access_token_detail.Token,
			RefershToken: *refresh_token_detail.Token,
		}, &Token{
			AccessToken:  access_token_detail,
			RefreshToken: refresh_token_detail,
		}, nil
}

func generateAccessToken(id int, username, tokenID string) (*TokenDetail, error) {
	// creating access claim
	access_claims := &Claims{
		Username: username,
		ID:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: AccessExpiresAt,
			IssuedAt:  IssuedAt,
			NotBefore: NotBefore,
			Subject:   "access_token",
			ID:        tokenID,
		},
	}

	// creating a new token with access claims
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)

	// sign the token with our unique secret from the environment files
	signedAccessToken, err := access_token.SignedString(secret)
	if err != nil {
		return nil, err
	}
	return &TokenDetail{
		Token:     &signedAccessToken,
		UserID:    id,
		Username:  username,
		ExpiresAt: access_claims.RegisteredClaims.ExpiresAt.Time,
		Subject:   access_claims.RegisteredClaims.Subject,
		TokenID:   access_claims.RegisteredClaims.ID,
	}, nil

}

func generateRefreshToken(id int, username, tokenID string) (*TokenDetail, error) {
	// creating custom refresh token claims
	refresh_claims := &Claims{
		Username: username,
		ID:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: RefreshExpiresAt,
			IssuedAt:  IssuedAt,
			NotBefore: NotBefore,
			Subject:   "refresh_token",
			ID:        tokenID,
		},
	}

	// creating a new token with refresh claims
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)

	// sign the token with the secret
	signedRefreshToken, err := refresh_token.SignedString(secret)
	if err != nil {
		return nil, err
	}
	return &TokenDetail{
		Token:     &signedRefreshToken,
		TokenID:   refresh_claims.RegisteredClaims.ID,
		UserID:    id,
		Username:  username,
		ExpiresAt: refresh_claims.RegisteredClaims.ExpiresAt.Time,
		Subject:   refresh_claims.RegisteredClaims.Subject,
	}, nil

}
