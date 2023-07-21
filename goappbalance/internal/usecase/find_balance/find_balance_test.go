package find_balance

import (
	"testing"
	"time"

	"github.com.br/rdo042/goappbalance/internal/entity"
	"github.com.br/rdo042/goappbalance/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFindBalanceUseCase_Execute(t *testing.T) {

	balance := entity.NewBalance("0276dc9e-1d30-41e9-8225-f2d251655f05", "4b9633e7-21bf-413c-be9c-4cded7468652", "c990b127-9c96-4ea9-a996-c70c3d9bd44c", 998.00, 2002.00, time.Now(), time.Now())

	balanceMock := &mocks.BalanceGatewayMock{}
	balanceMock.On("Get", balance.ID).Return(balance, nil)

	uc := NewFindBalanceUseCase(balanceMock)

	inputDto := FindBalanceInputDTO{
		BalanceID: balance.ID,
	}

	output, err := uc.Execute(inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	balanceMock.AssertExpectations(t)
	balanceMock.AssertNumberOfCalls(t, "Get", 1)
}
