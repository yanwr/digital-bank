package repositories

import (
	"errors"
	"yanwr/digital-bank/models"

	"gorm.io/gorm"
)

type IAccountRepository interface {
	FindByCpf(cpf string) (*models.Account, error)
	FindById(id string) (*models.Account, error)
	FindAll() ([]*models.Account, error)
	Create(account *models.Account) error
	Update(account *models.Account) error
}

type AccountRepository struct {
	connectionDB *gorm.DB
}

func NewAccountRepository(conDB *gorm.DB) IAccountRepository {
	return &AccountRepository{
		connectionDB: conDB,
	}
}

func (aR *AccountRepository) FindByCpf(cpf string) (*models.Account, error) {
	var account *models.Account
	aR.connectionDB.Find(&account, "cpf = ?", cpf)
	if len(account.Id) == 0 {
		return nil, errors.New("not found Account with cpf " + cpf)
	}
	return account, nil
}

func (aR *AccountRepository) FindById(id string) (*models.Account, error) {
	var account *models.Account
	aR.connectionDB.First(&account, "id = ?", id)
	if account == nil {
		return nil, errors.New("not found Account with id = " + id)
	}
	return account, nil
}

func (aR *AccountRepository) FindAll() ([]*models.Account, error) {
	var accounts []*models.Account
	aR.connectionDB.Find(&accounts)
	if accounts == nil {
		return nil, errors.New("error to find all accounts")
	}
	return accounts, nil
}

func (aR *AccountRepository) Create(account *models.Account) error {
	err := aR.connectionDB.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (aR *AccountRepository) Update(account *models.Account) error {
	err := aR.connectionDB.Save(account).Error
	if err != nil {
		return err
	}
	return nil
}
