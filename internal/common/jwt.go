package common

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT[T any](user *T) (string, error) {

	if user == nil {
		return "", errors.New("user can be null")
	}

	now := time.Now()

	exp := time.Hour * time.Duration(Expired())

	claims := &jwtUser[T]{
		User: user,
		Iat:  now.Unix(),
		Exp:  now.Add(exp).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(Secret())
}

func ParseWithExpiredJWT[T any](token string) (*T, error) {
	var сlaims jwtUser[T]

	jwtToken, err := jwt.ParseWithClaims(token, &сlaims, func(token *jwt.Token) (any, error) {
		return Secret(), nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, errors.New("invalid token")
	}

	return сlaims.User, nil
}

type jwtUser[T any] struct {
	User *T    `json:"user"`
	Iat  int64 `json:"iat"`
	Exp  int64 `json:"exp"`
}

func (j *jwtUser[T]) GetExpirationTime() (*jwt.NumericDate, error) {
	date := time.Unix(j.Exp, 0)
	return jwt.NewNumericDate(date), nil
}

func (j *jwtUser[T]) GetIssuedAt() (*jwt.NumericDate, error) {
	date := time.Unix(j.Iat, 0)
	return jwt.NewNumericDate(date), nil
}

// not used
func (j *jwtUser[T]) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

// not used
func (j *jwtUser[T]) GetIssuer() (string, error) {
	return "", nil
}

// not used
func (j *jwtUser[T]) GetSubject() (string, error) {
	return "", nil
}

// not used
func (j *jwtUser[T]) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}
