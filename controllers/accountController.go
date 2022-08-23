package controllers

import (
	"net/http"
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/exceptions"
	"yanwr/digital-bank/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IAccountController interface {
	IndexAllAccounts(c *gin.Context)
	ShowBalanceAccount(c *gin.Context)
	CreateAccount(c *gin.Context)
}

type AccountController struct {
	accountService services.IAccountService
}

func NewAccountController(conDB *gorm.DB) IAccountController {
	return &AccountController{
		accountService: services.NewAccountService(conDB),
	}
}

func (aC *AccountController) IndexAllAccounts(c *gin.Context) {
	accounts, err := aC.accountService.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func (aC *AccountController) ShowBalanceAccount(c *gin.Context) {
	accountId := c.Param("account_id")
	if len(accountId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, exceptions.ThrowBadRequestError("account_id can't be empty"))
		return
	}
	accountDto, errS := aC.accountService.FindById(string(accountId))
	if errS != nil {
		c.AbortWithStatusJSON(errS.Status, errS)
		return
	}
	c.JSON(http.StatusOK, accountDto.Balance)
}

func (aC *AccountController) CreateAccount(c *gin.Context) {
	var accountDto dtos.AccountRequestDTO
	err := c.BindJSON(&accountDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, exceptions.ThrowBadRequestError("invalid data to create a new account"))
		return
	}
	accountResponseDto, errS := aC.accountService.CreateAccount(accountDto)
	if errS != nil {
		c.AbortWithStatusJSON(errS.Status, errS)
		return
	}
	c.JSON(http.StatusCreated, accountResponseDto)
}
