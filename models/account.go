package models

import (
	"log"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var REGEX_CPF string = `[0-9]{3}\.?[0-9]{3}\.?[0-9]{3}\-?[0-9]{2}`

type Account struct {
	gorm.Model
	Id        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf" gorm:"unique"`
	Secret    string    `json:"secret"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func (account *Account) IsValid() error {
	err := IsCpfValid(account.Cpf)
	if err != nil {
		return err
	}
	return nil
}

func IsCpfValid(cpf string) error {
	match, err := regexp.MatchString(REGEX_CPF, cpf)
	if err != nil && !match {
		return err
	}
	return nil
}

func NewAccount(name string, cpf string, secret string) (*Account, error) {
	account := Account{
		Id:        uuid.New().String(),
		Name:      name,
		Cpf:       cpf,
		Secret:    hashSecret([]byte(secret)),
		Balance:   0,
		CreatedAt: time.Now(),
	}

	err := account.IsValid()
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func hashSecret(secret []byte) string {
	hash, err := bcrypt.GenerateFromPassword(secret, 8)
	if err != nil {
		log.Println(err)
		panic("failed to hash secret")
	}
	return string(hash)
}
