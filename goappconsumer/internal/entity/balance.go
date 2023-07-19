package entity

import (
	"errors"
	"time"
)

type Balance struct {
	AccountIdFrom        string
	AccountIdTo          string
	BalanceAccountIdFrom float64
	BalanceAccountIdTo   float64
}

func NewBalance(accountIdFrom string, accountIdTo string, balanceAccountIdFrom float64, balanceAccountIdTo float64) *Balance {

	balance := &Balance{
		AccountIdFrom:        accountIdFrom,
		AccountIdTo:          accountIdTo,
		balanceAccountIdFrom: balanceAccountIdFrom,
		balanceAccountIdTo:   balanceAccountIdTo,
	}

	return balance
}
