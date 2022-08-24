package services

import (
	"testing"
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/exceptions"
	"yanwr/digital-bank/mappers"
	"yanwr/digital-bank/models"
	"yanwr/digital-bank/tests/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWhenFindAllThenReturnTransfers(t *testing.T) {
	var (
		mockTransferRepository = new(mocks.MockTransferRepository)
		transferService        = &TransferService{transferRepository: mockTransferRepository, transferMapper: mappers.NewTransferMapper()}
	)
	t1 := models.NewTransfer("1hdsoie", "123", 500)
	t2 := models.NewTransfer("1hdsoie", "123", 800)
	expectedTransfers := []*models.Transfer{t1, t2}
	expectedTransfersDtos := transferService.transferMapper.ToDtos(expectedTransfers)

	mockTransferRepository.On("FindAllByAccountId", mock.Anything).Return(expectedTransfers, nil)
	transfersDtos, errS := transferService.FindAllByAccountId(t1.AccountOriginId)

	assert.Nil(t, errS)
	assert.NotEmpty(t, transfersDtos)
	assert.Equal(t, transfersDtos, expectedTransfersDtos)
}

func TestGivenEmptyAccountIdWhenFindAllThenReturnError(t *testing.T) {
	var (
		mockTransferRepository = new(mocks.MockTransferRepository)
		transferService        = &TransferService{transferRepository: mockTransferRepository, transferMapper: mappers.NewTransferMapper()}
	)
	expectedErrS := exceptions.ThrowBadRequestError("accountId can not be empty")
	transfersDtos, errS := transferService.FindAllByAccountId("")

	assert.Nil(t, transfersDtos)
	assert.NotNil(t, errS)
	assert.Equal(t, errS.Error_Message, expectedErrS.Error_Message)
}

func TestWhenCreateTransferToThenReturnTransfer(t *testing.T) {
	var (
		mockAccountRepository  = new(mocks.MockAccountRepository)
		mockTransferRepository = new(mocks.MockTransferRepository)
		mockAccountService     = new(mocks.MockAccountService)
		transferService        = &TransferService{
			transferRepository: mockTransferRepository,
			accountService:     &AccountService{accountRepository: mockAccountRepository, accountMapper: mappers.NewAccountMapper()},
			transferMapper:     mappers.NewTransferMapper(),
		}
	)
	accountOrigin, _ := models.NewAccount("John Estafano", "12345678940", "secret", 2500)
	accountDestination, _ := models.NewAccount("Sky Watson", "46537864589", "secret", 7000)
	accountOriginDto := dtos.NewAccountResponseDto(accountOrigin.Id, accountOrigin.Name, accountOrigin.Cpf, accountOrigin.Secret, accountOrigin.Balance, accountOrigin.CreatedAt)
	accountDestinationDto := dtos.NewAccountResponseDto(accountDestination.Id, accountDestination.Name, accountDestination.Cpf, accountDestination.Secret, accountDestination.Balance, accountDestination.CreatedAt)
	reqTransferDto := dtos.NewTransferRequestDto(accountDestination.Id, 500)
	expectedTransfer := models.NewTransfer(accountDestination.Id, accountOrigin.Id, 500)
	expectedTransferDto := dtos.NewTransferResponseDto(expectedTransfer.Id, expectedTransfer.AccountDestinationId, expectedTransfer.Amount, expectedTransfer.CreatedAt)

	mockTransferRepository.On("Create", mock.Anything).Return(nil)
	mockAccountRepository.On("FindById", accountOrigin.Id).Return(accountOrigin, nil)
	mockAccountRepository.On("FindById", accountDestination.Id).Return(accountDestination, nil)
	mockAccountRepository.On("Update", mock.Anything).Return(nil)

	mockAccountService.On("FindById", accountOrigin.Id).Return(accountOrigin, nil)
	mockAccountService.On("FindById", accountDestination.Id).Return(accountDestination, nil)
	mockAccountService.On("HasEnoughBalance", reqTransferDto.Amount, accountOrigin.Id).Return(nil)
	mockAccountService.On("UpdateAccount", accountOriginDto).Return(nil, nil)
	mockAccountService.On("UpdateAccount", accountDestinationDto).Return(nil, nil)

	resTransferDto, errS := transferService.CreateTransferTo(reqTransferDto, accountOrigin.Id)

	assert.Nil(t, errS)
	assert.Equal(t, resTransferDto.AccountDestinationId, expectedTransferDto.AccountDestinationId)
	assert.Equal(t, resTransferDto.Amount, expectedTransferDto.Amount)
}

func TestGivenAccountWithoutEnoughBalanceWhenCreateTransferToThenReturnError(t *testing.T) {
	var (
		mockAccountRepository  = new(mocks.MockAccountRepository)
		mockTransferRepository = new(mocks.MockTransferRepository)
		mockAccountService     = new(mocks.MockAccountService)
		transferService        = &TransferService{
			transferRepository: mockTransferRepository,
			accountService:     &AccountService{accountRepository: mockAccountRepository, accountMapper: mappers.NewAccountMapper()},
			transferMapper:     mappers.NewTransferMapper(),
		}
	)
	accountOrigin, _ := models.NewAccount("John Estafano", "12345678940", "secret", 2500)
	accountDestination, _ := models.NewAccount("Sky Watson", "46537864589", "secret", 7000)
	accountOriginDto := dtos.NewAccountResponseDto(accountOrigin.Id, accountOrigin.Name, accountOrigin.Cpf, accountOrigin.Secret, accountOrigin.Balance, accountOrigin.CreatedAt)
	accountDestinationDto := dtos.NewAccountResponseDto(accountDestination.Id, accountDestination.Name, accountDestination.Cpf, accountDestination.Secret, accountDestination.Balance, accountDestination.CreatedAt)
	reqTransferDto := dtos.NewTransferRequestDto(accountDestination.Id, 5000)
	expectedErroS := exceptions.ThrowBadRequestError("account origin does not have enough balance")

	mockTransferRepository.On("Create", mock.Anything).Return(nil)
	mockAccountRepository.On("FindById", accountOrigin.Id).Return(accountOrigin, nil)
	mockAccountRepository.On("FindById", accountDestination.Id).Return(accountDestination, nil)
	mockAccountRepository.On("Update", mock.Anything).Return(nil)

	mockAccountService.On("FindById", accountOrigin.Id).Return(accountOrigin, nil)
	mockAccountService.On("FindById", accountDestination.Id).Return(accountDestination, nil)
	mockAccountService.On("HasEnoughBalance", reqTransferDto.Amount, accountOrigin.Id).Return(nil)
	mockAccountService.On("UpdateAccount", accountOriginDto).Return(nil, nil)
	mockAccountService.On("UpdateAccount", accountDestinationDto).Return(nil, nil)

	resTransferDto, errS := transferService.CreateTransferTo(reqTransferDto, accountOrigin.Id)

	assert.Nil(t, resTransferDto)
	assert.NotNil(t, errS)
	assert.Equal(t, errS.Error_Message, expectedErroS.Error_Message)
}
