package middlewares

import (
	"strings"
	"yanwr/digital-bank/env"
	"yanwr/digital-bank/exceptions"
	"yanwr/digital-bank/services"

	"github.com/gin-gonic/gin"
)

type IAuthMiddleware interface {
	Authorize() gin.HandlerFunc
}

type AuthMiddleware struct {
	jwtService services.IJwtService
}

func NewAuthMiddleware() IAuthMiddleware {
	return &AuthMiddleware{
		jwtService: services.NewJWTService(),
	}
}

func (aM *AuthMiddleware) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(env.AUTHORIZATION_HEADER)
		if len(authHeader) == 0 {
			errS := exceptions.ThrowBadRequestError("token not found in Authorizarion Header")
			c.AbortWithStatusJSON(errS.Status, errS)
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			errS := exceptions.ThrowBadRequestError("invalid Authorization Header format")
			c.AbortWithStatusJSON(errS.Status, errS)
			return
		}

		authHeaderType := strings.ToUpper(fields[0])
		if env.AUTHORIZATION_HEADER_TYPE != authHeaderType {
			errS := exceptions.ThrowBadRequestError("unsupported Authorization Type " + authHeaderType)
			c.AbortWithStatusJSON(errS.Status, errS)
			return
		}

		tokenString := fields[1]
		payload, err := aM.jwtService.ValidateToken(tokenString)
		if err != nil {
			errS := exceptions.ThrowUnauthorizedError(err.Error())
			c.AbortWithStatusJSON(errS.Status, errS)
			return
		}

		c.Set(env.AUTHORIZATION_PAYLOAD, payload)
		c.Next()
	}
}
