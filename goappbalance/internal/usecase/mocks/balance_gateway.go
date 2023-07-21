package mocks

import (
	"github.com.br/rdo042/goappbalance/internal/entity"
	"github.com/stretchr/testify/mock"
)

type BalanceGatewayMock struct {
	mock.Mock
}

func (m *BalanceGatewayMock) Get(id string) (*entity.Balance, error) {

	args := m.Called(id)
	return args.Get(0).(*entity.Balance), args.Error(1)

}
