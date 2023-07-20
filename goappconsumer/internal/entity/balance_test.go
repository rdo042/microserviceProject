package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewBalance(t *testing.T) {
	balance, _ := NewBalance("11111111", "222222", 10, 20)

	assert.NotNil(t, balance)
	assert.Equal(t, balance.AccountIdFrom, "11111111")
	assert.Equal(t, balance.AccountIdTo, "222222")
	assert.Equal(t, balance.BalanceAccountIdFrom, 10.0)

}
