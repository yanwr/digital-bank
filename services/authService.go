package services

import (
	"log"
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/exceptions"
	"yanwr/digital-bank/mappers"
	"yanwr/digital-bank/models"
	"yanwr/digital-bank/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IAuthService interface {
	IsCredentialValid(cpf string, Secret string) (*dtos.AccountResponseDTO, *exceptions.StandardError)
}

type AuthSerivce struct {
	accountRepository repositories.IAccountRepository
	accountMapper     mappers.IAccountMapper
}

func NewAuthService(conDB *gorm.DB) IAuthService {
	return &AuthSerivce{
		accountRepository: repositories.NewAccountRepository(conDB),
		accountMapper:     mappers.NewAccountMapper(),
	}
}

func (aS *AuthSerivce) IsCredentialValid(cpf string, secret string) (*dtos.AccountResponseDTO, *exceptions.StandardError) {
	if err := models.IsCpfValid(cpf); err != nil {
		return nil, exceptions.ThrowBadRequestError("invalid CPF")
	}
	account, err := aS.accountRepository.FindByCpf(cpf)
	if err != nil {
		return nil, exceptions.ThrowInternalServerError("error to validade if there is account with the same CPF")
	}
	if len(account.Id) <= 0 {
		return nil, exceptions.ThrowBadRequestError("invalid credential, account does not exist")
	}
	comparedSecret := compareSecret(account.Secret, []byte(secret))
	if account.Cpf == cpf && comparedSecret {
		return aS.accountMapper.ToDto(account), nil
	}
	return nil, exceptions.ThrowBadRequestError("invalid secret, please try again")
}

func compareSecret(hashedPwd string, plainSecret []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainSecret)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
