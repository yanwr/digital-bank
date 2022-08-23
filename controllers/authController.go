package controllers

import (
	"net/http"
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/exceptions"
	"yanwr/digital-bank/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IAuthController interface {
	Login(c *gin.Context)
}

type AuthController struct {
	authService services.IAuthService
	jwtService  services.IJwtService
}

func NewAuthController(conDB *gorm.DB) IAuthController {
	return &AuthController{
		authService: services.NewAuthService(conDB),
		jwtService:  services.NewJWTService(),
	}
}

func (aC *AuthController) Login(c *gin.Context) {
	var login dtos.LoginRequestDTO
	err := c.BindJSON(&login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, exceptions.ThrowBadRequestError("invalid data to login"))
		return
	}
	accountDto, errS := aC.authService.IsCredentialValid(login.Cpf, login.Secret)
	if err != nil {
		c.AbortWithStatusJSON(errS.Status, errS)
		return
	}
	generatedToken := aC.jwtService.GenerateToken(accountDto.Id)
	c.JSON(http.StatusOK, dtos.LoginResponseDTO{Token: generatedToken})
}
