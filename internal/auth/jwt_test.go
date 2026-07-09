package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	// arrange
	userID := uuid.New()
	secret := "some-secret"
	expiresIn := time.Hour

	// act
	tokenString, err := MakeJWT(userID, secret, expiresIn)

	// asert part 1: making the token should not error
	if err != nil {
		t.Fatalf("expected no error making token: %v", err)
	}

	// act again
	gotUserID, err := ValidateJWT(tokenString, secret)

	// assert part 2
	if err != nil {
		t.Fatalf("expected no error validating token: %v", err)
	}

	if gotUserID != userID {
		t.Errorf("expected user ID %v, got %v", userID, gotUserID)
	}
}

func TestValidateJWTExpiredToken(t *testing.T) {
	userID := uuid.New()
	secret := "some-secret"

	tokenString, err := MakeJWT(userID, secret, -time.Hour)
	if err != nil {
		t.Fatalf("expected no error making token: %v", err)
	}

	_, err = ValidateJWT(tokenString, secret)
	if err == nil {
		t.Fatal("expected error validating expired token")
	}
}

func TestValidateJWTWrongSecret(t *testing.T) {
	userID := uuid.New()

	tokenString, err := MakeJWT(userID, "correct-secret", time.Hour)
	if err != nil {
		t.Fatalf("expected no error making token: %v", err)
	}

	_, err = ValidateJWT(tokenString, "wrong-secret")
	if err == nil {
		t.Fatal("expected error validating token with wrong secret")
	}
}
