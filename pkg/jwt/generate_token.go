package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func GenerateJwtToken(
	issuer string,
	secret string,
	subject string,
	expiresAfter time.Duration,
) (string, time.Time, error) {
	currentTime := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        uuid.New().String(),
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(currentTime),
		NotBefore: jwt.NewNumericDate(currentTime),
		ExpiresAt: jwt.NewNumericDate(currentTime.Add(expiresAfter)),
		Subject:   subject,
	})

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", currentTime, err
	}

	return signedToken, currentTime, nil
}
