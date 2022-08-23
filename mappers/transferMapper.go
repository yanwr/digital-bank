package mappers

import (
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/models"
)

type ITransferMapper interface {
	ToDtos(transfers []*models.Transfer) []*dtos.TransferResponseDTO
	ToDto(transfer *models.Transfer) *dtos.TransferResponseDTO
}

type TransferMapper struct{}

func NewTransferMapper() ITransferMapper {
	return &TransferMapper{}
}

func (tM *TransferMapper) ToDtos(transfers []*models.Transfer) []*dtos.TransferResponseDTO {
	var transfersDtos []*dtos.TransferResponseDTO
	for _, transfer := range transfers {
		transfersDtos = append(transfersDtos, tM.ToDto(transfer))
	}
	return transfersDtos
}

func (tM *TransferMapper) ToDto(transfer *models.Transfer) *dtos.TransferResponseDTO {
	return &dtos.TransferResponseDTO{
		Id:                   transfer.Id,
		AccountDestinationId: transfer.AccountDestinationId,
		Amount:               transfer.Amount,
		CreatedAt:            transfer.CreatedAt,
	}
}
