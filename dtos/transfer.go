package dtos

import "time"

type TransferRequestDTO struct {
	AccountDestinationId string  `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
}

func NewTransferRequestDto(account_destination_id string, amout float64) *TransferRequestDTO {
	return &TransferRequestDTO{
		AccountDestinationId: account_destination_id,
		Amount:               amout,
	}
}

type TransferResponseDTO struct {
	Id                   string    `json:"id" gorm:"primaryKey"`
	AccountDestinationId string    `json:"account_destination_id"`
	Amount               float64   `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

func NewTransferResponseDto(id string, account_destination_id string, amout float64, created_at time.Time) *TransferResponseDTO {
	return &TransferResponseDTO{
		Id:                   id,
		AccountDestinationId: account_destination_id,
		Amount:               amout,
		CreatedAt:            created_at,
	}
}
