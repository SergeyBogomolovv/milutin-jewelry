package utils

import "github.com/golang-jwt/jwt/v5"

func VerifyToken(token string, secret []byte) error {
	parser := jwt.NewParser(
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	parsed, err := parser.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !parsed.Valid {
		return jwt.ErrTokenSignatureInvalid
	}
	return nil
}
