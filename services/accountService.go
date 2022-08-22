package services

import (
	"log"
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
	ThereIsDuplicateAccounts(cpf string) *exceptions.StandardError
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

	accout, err := models.NewAccount(accountDto.Name, accountDto.Cpf, accountDto.Secret)
	if err != nil {
		return nil, exceptions.ThrowBadRequestError("invalid data to create new Account")
	}
	if err := aS.accountRepository.Create(accout); err != nil {
		log.Printf("Yan here 2 " + err.Error())
		return nil, exceptions.ThrowInternalServerError("error to create new Account")
	}
	return aS.accountMapper.ToDto(accout), nil
}

func (aS *AccountService) ThereIsDuplicateAccounts(cpf string) *exceptions.StandardError {
	if err := models.IsCpfValid(cpf); err != nil {
		return exceptions.ThrowBadRequestError("invalid CPF")
	}
	account, err := aS.accountRepository.FindByCpf(cpf)
	if err != nil {
		return exceptions.ThrowInternalServerError("error to validade if there is account with the same CPF")
	}
	if account != nil {
		return exceptions.ThrowBadRequestError("already exists account with the same CPF")
	}
	return nil
}
