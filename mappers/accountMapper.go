package mappers

import (
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/models"
)

type IAccountMapper interface {
	ToDtos(accounts []*models.Account) []*dtos.AccountResponseDTO
	ToDto(account *models.Account) *dtos.AccountResponseDTO
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
	var accountDto dtos.AccountResponseDTO

	accountDto.Balance = account.Balance
	accountDto.Cpf = account.Cpf
	accountDto.Name = account.Name
	accountDto.Id = account.Id
	accountDto.CreatedAt = account.CreatedAt
	return &accountDto
}
