package mocks

import (
	"yanwr/digital-bank/dtos"
	"yanwr/digital-bank/exceptions"

	"github.com/stretchr/testify/mock"
)

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) FindAll() ([]*dtos.AccountResponseDTO, *exceptions.StandardError) {
	ret := m.Called()

	var r0 []*dtos.AccountResponseDTO
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*dtos.AccountResponseDTO)
	}

	var r1 *exceptions.StandardError
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(*exceptions.StandardError)
	}

	return r0, r1
}

func (m *MockAccountService) FindById(id string) (*dtos.AccountResponseDTO, *exceptions.StandardError) {
	ret := m.Called(id)

	var r0 *dtos.AccountResponseDTO
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*dtos.AccountResponseDTO)
	}

	var r1 *exceptions.StandardError
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(*exceptions.StandardError)
	}

	return r0, r1
}

func (m *MockAccountService) CreateAccount(accountDto dtos.AccountRequestDTO) (*dtos.AccountResponseDTO, *exceptions.StandardError) {
	ret := m.Called(accountDto)

	var r0 *dtos.AccountResponseDTO
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*dtos.AccountResponseDTO)
	}

	var r1 *exceptions.StandardError
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(*exceptions.StandardError)
	}

	return r0, r1
}

func (m *MockAccountService) UpdateAccount(accountDto *dtos.AccountResponseDTO) (*dtos.AccountResponseDTO, *exceptions.StandardError) {
	ret := m.Called(accountDto)

	var r0 *dtos.AccountResponseDTO
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*dtos.AccountResponseDTO)
	}

	var r1 *exceptions.StandardError
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(*exceptions.StandardError)
	}

	return r0, r1
}

func (m *MockAccountService) ThereIsDuplicateAccounts(cpf string) *exceptions.StandardError {
	ret := m.Called(cpf)

	var r1 *exceptions.StandardError
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(*exceptions.StandardError)
	}

	return r1
}

func (m *MockAccountService) HasEnoughBalance(balance float64, accountId string) *exceptions.StandardError {
	ret := m.Called(balance, accountId)

	var r1 *exceptions.StandardError
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(*exceptions.StandardError)
	}

	return r1
}
