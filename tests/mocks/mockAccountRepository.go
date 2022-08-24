package mocks

import (
	"yanwr/digital-bank/models"

	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) FindByCpf(cpf string) (*models.Account, error) {
	ret := m.Called(cpf)

	var r0 *models.Account
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.Account)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockAccountRepository) FindById(id string) (*models.Account, error) {
	ret := m.Called(id)

	var r0 *models.Account
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.Account)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockAccountRepository) FindAll() ([]*models.Account, error) {
	ret := m.Called()

	var r0 []*models.Account
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*models.Account)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockAccountRepository) Create(account *models.Account) error {
	ret := m.Called(account)

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r1
}

func (m *MockAccountRepository) Update(account *models.Account) error {
	ret := m.Called(account)

	var r1 error
	if ret.Get(0) != nil {
		r1 = ret.Get(0).(error)
	}

	return r1
}
