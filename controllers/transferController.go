package controllers

import (
	"net/http"
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/env"
	"yanwr/digital-bank/exceptions"
	"yanwr/digital-bank/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ITransferController interface {
	IndexAllTransfersFromCurrentUser(c *gin.Context)
	CreateTransferTo(c *gin.Context)
}

type TransferController struct {
	transferService services.ITransferService
}

func NewTransferController(conDB *gorm.DB) ITransferController {
	return &TransferController{
		transferService: services.NewTransferService(conDB),
	}
}

func (tC *TransferController) IndexAllTransfersFromCurrentUser(c *gin.Context) {
	payload := c.MustGet(env.AUTHORIZATION_PAYLOAD).(*dtos.Payload)
	transfers, errS := tC.transferService.FindAllByAccountId(payload.AccountId)
	if errS != nil {
		c.AbortWithStatusJSON(errS.Status, errS)
		return
	}
	c.JSON(http.StatusOK, transfers)
}

func (tC *TransferController) CreateTransferTo(c *gin.Context) {
	var transferDtoReq *dtos.TransferRequestDTO
	err := c.BindJSON(&transferDtoReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, exceptions.ThrowBadRequestError("invalid data to create a transfer"))
		return
	}
	paylod := c.MustGet(env.AUTHORIZATION_PAYLOAD).(*dtos.Payload)
	transferDtoRes, errS := tC.transferService.CreateTransferTo(transferDtoReq, paylod.AccountId)
	if errS != nil {
		c.AbortWithStatusJSON(errS.Status, errS)
		return
	}
	c.JSON(http.StatusOK, transferDtoRes)
}
