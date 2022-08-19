package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	Id                   string    `json:"id" gorm:"primaryKey" valid:"required"`
	AccountOriginId      int       `json:"account_origin_id" valid:"notnull"`
	AccountDestinationId int       `json:"account_destination_id" valid:"notnull"`
	Amount               float64   `json:"amount" valid:"notnull"`
	CreatedAt            time.Time `json:"created_at" valid:"required"`
	UpdatedAt            time.Time `json:"updated_at" valid:"required"`
	DeletedAt            time.Time `json:"deleted_at" valid:"required"`
}

func (transfer *Transfer) isValid() error {
	_, err := govalidator.ValidateStruct(transfer)
	if err != nil {
		return err
	}
	return nil
}
