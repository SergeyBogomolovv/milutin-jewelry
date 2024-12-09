package utils

import "github.com/golang-jwt/jwt/v5"

func VerifyToken(token string, secret []byte) error {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})
	if err != nil || !parsed.Valid {
		return jwt.ErrTokenNotValidYet
	}
	return nil
}
