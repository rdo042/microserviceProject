package entity

import (
	"time"
)

type Balance struct {
	ID                   string
	AccountIdFrom        string
	AccountIdTo          string
	BalanceAccountIdFrom float64
	BalanceAccountIdTo   float64
	CreatedAt            time.Time
	UpdateAt             time.Time
}

func NewBalance(id string, accountIdFrom string, accountIdTo string, balanceAccountIdFrom float64, balanceAccountIdTo float64, createdAt time.Time, updateAt time.Time) *Balance {

	balance := &Balance{
		ID:                   id,
		AccountIdFrom:        accountIdFrom,
		AccountIdTo:          accountIdTo,
		BalanceAccountIdFrom: balanceAccountIdFrom,
		BalanceAccountIdTo:   balanceAccountIdTo,
		CreatedAt:            createdAt,
		UpdateAt:             updateAt,
	}

	return balance
}
