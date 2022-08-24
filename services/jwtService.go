package services

import (
	"errors"
	"os"
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/env"

	"github.com/dgrijalva/jwt-go"
)

type IJwtService interface {
	GenerateToken(accountId string) string
	ValidateToken(tokenString string) (*dtos.Payload, error)
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
	payload, err := dtos.NewPayload(accountId)
	if err != nil {
		panic(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(jS.secretKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func (jS *JwtService) ValidateToken(tokenString string) (*dtos.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(jS.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &dtos.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errors.New("token has expired")) {
			return nil, errors.New("token has expired")
		}
		return nil, errors.New("invalid token")
	}

	payload, ok := jwtToken.Claims.(*dtos.Payload)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return payload, nil
}
