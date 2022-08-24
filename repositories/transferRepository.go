package repositories

import (
	"fmt"
	"yanwr/digital-bank/models"

	"gorm.io/gorm"
)

type ITransferRepository interface {
	FindAllByAccountId(accountId string) ([]*models.Transfer, error)
	Create(transfer *models.Transfer) error
}

type TransferRepository struct {
	connectionDB *gorm.DB
}

func NewTransferRepository(conDB *gorm.DB) ITransferRepository {
	return &TransferRepository{
		connectionDB: conDB,
	}
}

func (tR *TransferRepository) FindAllByAccountId(accountId string) ([]*models.Transfer, error) {
	var transfers []*models.Transfer
	if err := tR.connectionDB.Find(&transfers, "account_origin_id = ?", accountId).Error; err != nil {
		return nil, fmt.Errorf("error to find all transfers by accountId = %s", accountId)
	}
	return transfers, nil
}

func (tR *TransferRepository) Create(transfer *models.Transfer) error {
	err := tR.connectionDB.Create(transfer).Error
	if err != nil {
		return err
	}
	return nil
}
