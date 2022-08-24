package services

import (
	"fmt"
	"testing"
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/exceptions"
	"yanwr/digital-bank/mappers"
	"yanwr/digital-bank/models"
	"yanwr/digital-bank/tests/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWhenFindAllThenReturnAccounts(t *testing.T) {
	var (
		mockAccountRepository = new(mocks.MockAccountRepository)
		accountService        = &AccountService{accountRepository: mockAccountRepository, accountMapper: mappers.NewAccountMapper()}
	)
	a1, _ := models.NewAccount("John Estafano", "12345678940", "secret", 2500)
	a2, _ := models.NewAccount("Sky Watson", "46537864589", "secret", 7000)
	expectedAccounts := []*models.Account{a1, a2}
	expectedAccountsDtos := accountService.accountMapper.ToDtos(expectedAccounts)

	mockAccountRepository.On("FindAll").Return(expectedAccounts, nil)
	accountsDtos, errS := accountService.FindAll()

	assert.Nil(t, errS)
	assert.NotEmpty(t, accountsDtos)
	assert.Equal(t, accountsDtos, expectedAccountsDtos)
}

func TestGivenValidIdWhenFindByIdThenReturnAccount(t *testing.T) {
	var (
		mockAccountRepository = new(mocks.MockAccountRepository)
		accountService        = &AccountService{accountRepository: mockAccountRepository, accountMapper: mappers.NewAccountMapper()}
	)
	a1, _ := models.NewAccount("John Estafano", "12345678940", "secret", 2500)
	expectedAccountDto := accountService.accountMapper.ToDto(a1)

	mockAccountRepository.On("FindById", mock.Anything).Return(a1, nil)
	accountsDto, errS := accountService.FindById(a1.Id)

	assert.Nil(t, errS)
	assert.NotNil(t, accountsDto)
	assert.Equal(t, accountsDto, expectedAccountDto)
}

func TestGivenInvalidIdWhenFindByIdThenReturnError(t *testing.T) {
	var (
		mockAccountRepository = new(mocks.MockAccountRepository)
		accountService        = &AccountService{accountRepository: mockAccountRepository, accountMapper: mappers.NewAccountMapper()}
	)
	invalidId := "123-invalid-id"
	expectedErrS := exceptions.ThrowNotFoundError("not found Account")

	mockAccountRepository.On("FindById", mock.Anything).Return(nil, fmt.Errorf("not found Account with id = %s", invalidId))
	accountDto, errS := accountService.FindById(invalidId)

	assert.Nil(t, accountDto)
	assert.NotNil(t, errS)
	assert.Equal(t, errS.Status, expectedErrS.Status)
	assert.Equal(t, errS.Error_Message, expectedErrS.Error_Message)
}

func TestGivenRequestAccountValidWhenCreateThenReturnOK(t *testing.T) {
	var (
		mockAccountRepository = new(mocks.MockAccountRepository)
		accountService        = &AccountService{accountRepository: mockAccountRepository, accountMapper: mappers.NewAccountMapper()}
	)
	reqAccountDto := *dtos.NewAccountRequestDto("John Estafano", "12345678940", "secret", 2500)
	expectedAccount, _ := models.NewAccount(reqAccountDto.Name, reqAccountDto.Cpf, reqAccountDto.Secret, reqAccountDto.Balance)
	mockAccountRepository.On("FindByCpf", mock.Anything).Return(nil, nil)
	mockAccountRepository.On("Create", mock.Anything).Return(expectedAccount, nil)
	accountsDto, errS := accountService.CreateAccount(reqAccountDto)

	assert.Nil(t, errS)
	assert.NotNil(t, accountsDto)
	assert.NotNil(t, accountsDto.Id)
}

func TestGivenRequestDuplicateAccountWhenCreateThenReturnError(t *testing.T) {
	var (
		mockAccountRepository = new(mocks.MockAccountRepository)
		accountService        = &AccountService{accountRepository: mockAccountRepository, accountMapper: mappers.NewAccountMapper()}
	)
	reqAccountDto := *dtos.NewAccountRequestDto("John Estafano", "12345678940", "secret", 2500)
	expectedAccount, _ := models.NewAccount(reqAccountDto.Name, reqAccountDto.Cpf, reqAccountDto.Secret, reqAccountDto.Balance)
	expectedErrS := exceptions.ThrowBadRequestError("already exists account with the same CPF")
	mockAccountRepository.On("FindByCpf", mock.Anything).Return(expectedAccount, nil)
	accountsDto, errS := accountService.CreateAccount(reqAccountDto)

	assert.Nil(t, accountsDto)
	assert.NotNil(t, errS)
	assert.Equal(t, errS.Status, expectedErrS.Status)
	assert.Equal(t, errS.Error_Message, expectedErrS.Error_Message)
}

func TestGivenAccountWhenUpdateThenReturnOK(t *testing.T) {
	var (
		mockAccountRepository = new(mocks.MockAccountRepository)
		accountService        = &AccountService{accountRepository: mockAccountRepository, accountMapper: mappers.NewAccountMapper()}
	)
	reqAccountDto := *dtos.NewAccountRequestDto("John Estafano", "12345678940", "secret", 2500)
	expectedAccount, _ := models.NewAccount(reqAccountDto.Name, reqAccountDto.Cpf, reqAccountDto.Secret, reqAccountDto.Balance)
	oldBalance := expectedAccount.Balance

	expectedAccount.Balance = 1000
	mockAccountRepository.On("Update", mock.Anything).Return(nil, nil)
	accountsDto, errS := accountService.UpdateAccount(accountService.accountMapper.ToDto(expectedAccount))

	assert.Nil(t, errS)
	assert.NotNil(t, accountsDto)
	assert.NotEqual(t, oldBalance, accountsDto.Balance)
}

func TestGivenEnoughBalanceWhenVerifyIfHasBalanceThenReturnOK(t *testing.T) {
	var (
		mockAccountRepository = new(mocks.MockAccountRepository)
		accountService        = &AccountService{accountRepository: mockAccountRepository, accountMapper: mappers.NewAccountMapper()}
	)
	reqAccountDto := *dtos.NewAccountRequestDto("John Estafano", "12345678940", "secret", 2500)
	expectedAccount, _ := models.NewAccount(reqAccountDto.Name, reqAccountDto.Cpf, reqAccountDto.Secret, reqAccountDto.Balance)

	mockAccountRepository.On("FindById", mock.Anything).Return(expectedAccount, nil)
	errS := accountService.HasEnoughBalance(1000, expectedAccount.Id)

	assert.Nil(t, errS)
}

func TestGivenNotEnoughBalanceWhenVerifyIfHasBalanceThenReturnError(t *testing.T) {
	var (
		mockAccountRepository = new(mocks.MockAccountRepository)
		accountService        = &AccountService{accountRepository: mockAccountRepository, accountMapper: mappers.NewAccountMapper()}
	)
	reqAccountDto := *dtos.NewAccountRequestDto("John Estafano", "12345678940", "secret", 2500)
	expectedAccount, _ := models.NewAccount(reqAccountDto.Name, reqAccountDto.Cpf, reqAccountDto.Secret, reqAccountDto.Balance)
	expectedErrS := exceptions.ThrowBadRequestError("account origin does not have enough balance")

	mockAccountRepository.On("FindById", mock.Anything).Return(expectedAccount, nil)
	errS := accountService.HasEnoughBalance(4000, expectedAccount.Id)

	assert.NotNil(t, errS)
	assert.Equal(t, errS.Error_Message, expectedErrS.Error_Message)
}
