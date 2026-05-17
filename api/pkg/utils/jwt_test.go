package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestVerifyTokenRequiresExpiration(t *testing.T) {
	token := signTestToken(t, jwt.RegisteredClaims{})

	if err := VerifyToken(token, []byte("secret")); err == nil {
		t.Fatal("expected token without expiration to be rejected")
	}
}

func TestVerifyTokenRejectsExpiredToken(t *testing.T) {
	token := signTestToken(t, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Minute)),
	})

	if err := VerifyToken(token, []byte("secret")); err == nil {
		t.Fatal("expected expired token to be rejected")
	}
}

func TestVerifyTokenAcceptsValidToken(t *testing.T) {
	token := signTestToken(t, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
	})

	if err := VerifyToken(token, []byte("secret")); err != nil {
		t.Fatalf("expected token to be valid: %v", err)
	}
}

func signTestToken(t *testing.T, claims jwt.RegisteredClaims) string {
	t.Helper()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}
	return token
}
