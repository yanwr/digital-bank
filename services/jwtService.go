package services

import (
	"fmt"
	"os"
	"time"
	"yanwr/digital-bank/env"

	"github.com/dgrijalva/jwt-go"
)

type IJwtService interface {
	GenerateToken(accountId string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type JwtService struct {
	secretKey string
}

func NewJWTService() IJwtService {
	return &JwtService{
		secretKey: os.Getenv(env.JWT_SECRET),
	}
}

func (jS *JwtService) GenerateToken(accountId string) string {
	claims := &jwt.StandardClaims{
		Id:        accountId,
		ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jS.secretKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func (jS *JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(jS.secretKey), nil
	})
}
