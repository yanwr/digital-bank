package models

import (
	"errors"
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
	Id        string    `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name,omitempty"`
	Cpf       string    `json:"cpf,omitempty" gorm:"unique"`
	Secret    string    `json:"secret,omitempty"`
	Balance   float64   `json:"balance,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func NewAccount(name string, cpf string, secret string, balance float64) (*Account, error) {
	account := Account{
		Id:        uuid.New().String(),
		Name:      name,
		Cpf:       cpf,
		Secret:    hashSecret([]byte(secret)),
		Balance:   balance,
		CreatedAt: time.Now(),
	}

	err := account.IsValid()
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (account *Account) IsValid() error {
	err := IsCpfValid(account.Cpf)
	if err != nil {
		return err
	}
	return nil
}

func IsCpfValid(cpf string) error {
	if len(cpf) != 11 {
		return errors.New("invalid format CPF, try: 'xxxxxxxxxxx'")
	}
	rgx, err := regexp.Compile(REGEX_CPF)
	if err != nil {
		return errors.New("error to compile Regex")
	}
	if !rgx.MatchString(cpf) {
		return errors.New("invalid CPF")
	}
	return nil
}

func hashSecret(secret []byte) string {
	hash, err := bcrypt.GenerateFromPassword(secret, 8)
	if err != nil {
		log.Println(err)
		panic("failed to hash secret")
	}
	return string(hash)
}
