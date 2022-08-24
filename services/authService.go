package services

import (
	"log"
	"yanwr/digital-bank/exceptions"
	"yanwr/digital-bank/mappers"
	"yanwr/digital-bank/models"
	"yanwr/digital-bank/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IAuthService interface {
	IsCredentialValid(cpf string, secret string) (string, *exceptions.StandardError)
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

func (aS *AuthSerivce) IsCredentialValid(cpf string, secret string) (string, *exceptions.StandardError) {
	err := models.IsCpfValid(cpf)
	if err != nil {
		return "", exceptions.ThrowBadRequestError(err.Error())
	}
	account, err := aS.accountRepository.FindByCpf(cpf)
	if err != nil {
		return "", exceptions.ThrowNotFoundError(err.Error())
	}
	comparedSecret := compareSecret(account.Secret, []byte(secret))
	if account.Cpf == cpf && comparedSecret {
		return account.Id, nil
	}

	return "", exceptions.ThrowBadRequestError("invalid secret, please try again")
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
