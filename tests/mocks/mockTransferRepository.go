package mocks

import (
	"yanwr/digital-bank/models"

	"github.com/stretchr/testify/mock"
)

type MockTransferRepository struct {
	mock.Mock
}

func (m *MockTransferRepository) FindAllByAccountId(accountId string) ([]*models.Transfer, error) {
	ret := m.Called(accountId)

	var r0 []*models.Transfer
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*models.Transfer)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockTransferRepository) Create(transfer *models.Transfer) error {
	ret := m.Called(transfer)

	var r1 error
	if ret.Get(0) != nil {
		r1 = ret.Get(0).(error)
	}

	return r1
}
