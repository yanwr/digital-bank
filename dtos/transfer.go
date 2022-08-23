package dtos

import "time"

type TransferRequestDTO struct {
	AccountDestinationId string  `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
}

type TransferResponseDTO struct {
	Id                   string    `json:"id" gorm:"primaryKey"`
	AccountDestinationId string    `json:"account_destination_id"`
	Amount               float64   `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}
