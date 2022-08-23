package services

import (
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/exceptions"
	"yanwr/digital-bank/mappers"
	"yanwr/digital-bank/models"
	"yanwr/digital-bank/repositories"

	"gorm.io/gorm"
)

type IAccountService interface {
	FindAll() ([]*dtos.AccountResponseDTO, *exceptions.StandardError)
	FindById(id string) (*dtos.AccountResponseDTO, *exceptions.StandardError)
	CreateAccount(accountDto dtos.AccountRequestDTO) (*dtos.AccountResponseDTO, *exceptions.StandardError)
	UpdateAccount(accountDto *dtos.AccountResponseDTO) (*dtos.AccountResponseDTO, *exceptions.StandardError)
	ThereIsDuplicateAccounts(cpf string) *exceptions.StandardError
	HasEnoughBalance(balance float64, accountId string) *exceptions.StandardError
}

type AccountService struct {
	accountRepository repositories.IAccountRepository
	accountMapper     mappers.IAccountMapper
}

func NewAccountService(conDB *gorm.DB) IAccountService {
	return &AccountService{
		accountRepository: repositories.NewAccountRepository(conDB),
		accountMapper:     mappers.NewAccountMapper(),
	}
}

func (aS *AccountService) FindAll() ([]*dtos.AccountResponseDTO, *exceptions.StandardError) {
	accounts, err := aS.accountRepository.FindAll()
	if err != nil {
		return nil, exceptions.ThrowInternalServerError("error to gel all Accounts")
	}
	return aS.accountMapper.ToDtos(accounts), nil
}

func (aS *AccountService) FindById(id string) (*dtos.AccountResponseDTO, *exceptions.StandardError) {
	account, err := aS.accountRepository.FindById(id)
	if err != nil {
		return nil, exceptions.ThrowNotFoundError("not found Account")
	}
	return aS.accountMapper.ToDto(account), nil
}

func (aS *AccountService) CreateAccount(accountDto dtos.AccountRequestDTO) (*dtos.AccountResponseDTO, *exceptions.StandardError) {
	if errS := aS.ThereIsDuplicateAccounts(accountDto.Cpf); errS != nil {
		return nil, errS
	}

	accout, err := models.NewAccount(accountDto.Name, accountDto.Cpf, accountDto.Secret, accountDto.Balance)
	if err != nil {
		return nil, exceptions.ThrowBadRequestError("invalid data to create new Account")
	}
	if err := aS.accountRepository.Create(accout); err != nil {
		return nil, exceptions.ThrowInternalServerError("error to create new Account")
	}
	return aS.accountMapper.ToDto(accout), nil
}

func (aS *AccountService) UpdateAccount(accountDto *dtos.AccountResponseDTO) (*dtos.AccountResponseDTO, *exceptions.StandardError) {
	account := aS.accountMapper.ToEntity(accountDto)
	if err := aS.accountRepository.Update(account); err != nil {
		return nil, exceptions.ThrowInternalServerError("error to update account")
	}
	return accountDto, nil
}

func (aS *AccountService) ThereIsDuplicateAccounts(cpf string) *exceptions.StandardError {
	err := models.IsCpfValid(cpf)
	if err != nil {
		return exceptions.ThrowBadRequestError(err.Error())
	}
	account, err := aS.accountRepository.FindByCpf(cpf)
	if err != nil {
		return exceptions.ThrowInternalServerError("error to validade if there is account with the same CPF")
	}
	if len(account.Id) > 0 {
		return exceptions.ThrowBadRequestError("already exists account with the same CPF")
	}
	return nil
}

func (aS *AccountService) HasEnoughBalance(balance float64, accountId string) *exceptions.StandardError {
	account, errS := aS.FindById(accountId)
	if errS != nil {
		return errS
	}

	if account.Balance >= balance {
		return nil
	}
	return exceptions.ThrowBadRequestError("account origin does not have enough balance")
}
