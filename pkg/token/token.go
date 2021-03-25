package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/leminhson2398/zipper/modules/setting"
	"github.com/leminhson2398/zipper/pkg/logger"
)

// TokenClaims is the claims the access JWT token contains
type TokenClaims struct {
	jwt.StandardClaims
}

// ErrMalformedToken is the error returned if the token has malformed
type ErrMalformedToken struct{}

// Error returns the error message for ErrMalformedToken
func (r *ErrMalformedToken) Error() string {
	return "token is malformed"
}

// ErrExpiredToken is the error returned if the token has expired
type ErrExpiredToken struct{}

// Error returns the error message for ErrExpiredToken
func (r *ErrExpiredToken) Error() string {
	return "token is expired"
}

// NewAccessToken generates a new JWT access token with the correct claims
func NewAccessToken(duration time.Duration) (string, error) {
	tokenExpirationTime := time.Now().Add(duration).Unix()
	tokenClaims := &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpirationTime,
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(setting.SecretSalt))
	if err != nil {
		return "", err
	}
	return accessTokenString, nil
}

// ValidateAccessToken validates a JWT access token and returns the contained claims or an error if it's invalid
func ValidateAccessToken(token string) (TokenClaims, error) {
	accessClaims := &TokenClaims{}
	accessToken, err := jwt.ParseWithClaims(token, accessClaims, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(setting.SecretSalt), nil
	})

	if accessToken.Valid {
		logger.Logger.Debug().Msg("token is valid")
		return *accessClaims, nil
	}

	if validationErr, ok := err.(*jwt.ValidationError); ok {
		if validationErr.Errors&(jwt.ValidationErrorMalformed|jwt.ValidationErrorSignatureInvalid) != 0 {
			return TokenClaims{}, &ErrMalformedToken{}
		} else if validationErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return TokenClaims{}, &ErrExpiredToken{}
		}
	}

	return TokenClaims{}, err
}
