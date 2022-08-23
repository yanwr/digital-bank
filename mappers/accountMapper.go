package mappers

import (
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/models"
)

type IAccountMapper interface {
	ToDtos(accounts []*models.Account) []*dtos.AccountResponseDTO
	ToDto(account *models.Account) *dtos.AccountResponseDTO
	ToEntity(accountDto *dtos.AccountResponseDTO) *models.Account
}

type AccountMapper struct{}

func NewAccountMapper() IAccountMapper {
	return &AccountMapper{}
}

func (aM *AccountMapper) ToDtos(accounts []*models.Account) []*dtos.AccountResponseDTO {
	var accountsDtos []*dtos.AccountResponseDTO
	for _, account := range accounts {
		accountsDtos = append(accountsDtos, aM.ToDto(account))
	}
	return accountsDtos
}

func (aM *AccountMapper) ToDto(account *models.Account) *dtos.AccountResponseDTO {
	return &dtos.AccountResponseDTO{
		Id:        account.Id,
		Cpf:       account.Cpf,
		Name:      account.Name,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}
}

func (aM *AccountMapper) ToEntity(accountDto *dtos.AccountResponseDTO) *models.Account {
	return &models.Account{
		Id:        accountDto.Id,
		Cpf:       accountDto.Cpf,
		Name:      accountDto.Name,
		Balance:   accountDto.Balance,
		CreatedAt: accountDto.CreatedAt,
	}
}
