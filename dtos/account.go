package dtos

import (
	"time"
)

type AccountRequestDTO struct {
	Name    string  `json:"name"`
	Cpf     string  `json:"cpf"`
	Secret  string  `json:"secret"`
	Balance float64 `json:"balance"`
}

func NewAccountRequestDto(name string, cpf string, secret string, balance float64) *AccountRequestDTO {
	return &AccountRequestDTO{
		Name:    name,
		Cpf:     cpf,
		Secret:  secret,
		Balance: balance,
	}
}

type AccountResponseDTO struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccountResponseDto(id string, name string, cpf string, secret string, balance float64, created_at time.Time) *AccountResponseDTO {
	return &AccountResponseDTO{
		Id:        id,
		Name:      name,
		Cpf:       cpf,
		Balance:   balance,
		CreatedAt: created_at,
	}
}
