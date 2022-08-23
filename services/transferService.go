package services

import (
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/exceptions"
	"yanwr/digital-bank/mappers"
	"yanwr/digital-bank/models"
	"yanwr/digital-bank/repositories"

	"gorm.io/gorm"
)

type ITransferService interface {
	FindAllByAccountId(accountId string) ([]*dtos.TransferResponseDTO, *exceptions.StandardError)
	CreateTransferTo(transferDto dtos.TransferRequestDTO, accountOriginId string) (*dtos.TransferResponseDTO, *exceptions.StandardError)
	updateBalance(accountOriginId string, accountDestinationId string, amount float64) *exceptions.StandardError
}

type TransferService struct {
	transferRepository repositories.ITransferRepository
	accountService     IAccountService
	transferMapper     mappers.ITransferMapper
}

func NewTransferService(conDB *gorm.DB) ITransferService {
	return &TransferService{
		transferRepository: repositories.NewTransferRepository(conDB),
		accountService:     NewAccountService(conDB),
		transferMapper:     mappers.NewTransferMapper(),
	}
}

func (tS *TransferService) FindAllByAccountId(accountId string) ([]*dtos.TransferResponseDTO, *exceptions.StandardError) {
	if len(accountId) <= 0 {
		return nil, exceptions.ThrowBadRequestError("accountId can not be empty")
	}
	transfers, err := tS.transferRepository.FindAllByAccountId(accountId)
	if err != nil {
		return nil, exceptions.ThrowInternalServerError(err.Error())
	}
	return tS.transferMapper.ToDtos(transfers), nil
}

func (tS *TransferService) CreateTransferTo(transferDto dtos.TransferRequestDTO, accountOriginId string) (*dtos.TransferResponseDTO, *exceptions.StandardError) {
	errS := tS.updateBalance(accountOriginId, transferDto.AccountDestinationId, transferDto.Amount)
	if errS != nil {
		return nil, errS
	}

	transfer := models.NewTransfer(transferDto.AccountDestinationId, accountOriginId, transferDto.Amount)
	err := tS.transferRepository.Create(transfer)
	if err != nil {
		return nil, exceptions.ThrowInternalServerError(err.Error())
	}
	return tS.transferMapper.ToDto(transfer), nil
}

func (tS *TransferService) updateBalance(accountOriginId string, accountDestinationId string, amount float64) *exceptions.StandardError {
	errS := tS.accountService.HasEnoughBalance(amount, accountOriginId)
	if errS != nil {
		return errS
	}

	accountOrigin, errS := tS.accountService.FindById(accountOriginId)
	if errS != nil {
		return errS
	}

	accountDestination, errS := tS.accountService.FindById(accountDestinationId)
	if errS != nil {
		return errS
	}

	accountOrigin.Balance = accountOrigin.Balance - amount
	accountDestination.Balance = accountDestination.Balance + amount

	tS.accountService.UpdateAccount(accountOrigin)
	tS.accountService.UpdateAccount(accountDestination)
	return nil
}
