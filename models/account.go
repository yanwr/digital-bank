package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Id        string    `json:"id" gorm:"primaryKey" valid:"required"`
	Name      string    `json:"name" valid:"notnull"`
	Cpf       string    `json:"cpf" valid:"notnull"`
	Secret    string    `json:"secret" valid:"notnull"`
	Balance   float64   `json:"balance" valid:"notnull"`
	CreatedAt time.Time `json:"created_at" valid:"required"`
	UpdatedAt time.Time `json:"updated_at" valid:"required"`
	DeletedAt time.Time `json:"deleted_at" valid:"required"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}
