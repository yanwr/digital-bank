package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	Id                   string    `json:"id,omitempty" gorm:"type:uuid;primaryKey"`
	AccountOriginId      string    `json:"account_origin_id,omitempty"`
	AccountDestinationId string    `json:"account_destination_id,omitempty"`
	Amount               float64   `json:"amount,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
}

func NewTransfer(accountDestinationId string, accountOriginId string, amout float64) *Transfer {
	return &Transfer{
		Id:                   uuid.New().String(),
		AccountOriginId:      accountOriginId,
		AccountDestinationId: accountDestinationId,
		Amount:               amout,
		CreatedAt:            time.Now(),
	}
}
