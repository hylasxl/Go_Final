package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("jwtsecretkeyahihi")

type JWTClaims struct {
	Username    string    `json:"username"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Address     string    `json:"address"`
	RoleID      uint      `json:"role_id"`
	IsActive    bool      `json:"is_active"`
	jwt.StandardClaims
}

func GenerateAccessToken(claims JWTClaims) (string, error) {
	claims.ExpiresAt = time.Now().Add(15 * time.Minute).Unix()
	claims.IssuedAt = time.Now().Unix()
	claims.Subject = claims.Username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func GenerateRefreshToken(username string) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
	}
	return claims, nil
}
