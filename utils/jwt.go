package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type JWTClaim struct {
	Wallet string `json:"wallet"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(wallet string, role string) (string, error) {
	claims := &JWTClaim{
		Wallet: wallet,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
