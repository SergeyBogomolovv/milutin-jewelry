package authusecase

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (u *usecase) signToken() (string, error) {
	iat := time.Now()
	exp := iat.Add(24 * time.Hour)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(iat),
		ExpiresAt: jwt.NewNumericDate(exp),
	}).SignedString(u.jwtSecret)
}
