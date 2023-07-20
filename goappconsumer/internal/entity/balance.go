package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
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

func NewBalance(accountIdFrom string, accountIdTo string, balanceAccountIdFrom float64, balanceAccountIdTo float64) (*Balance, error) {

	balance := &Balance{
		ID:                   uuid.New().String(),
		AccountIdFrom:        accountIdFrom,
		AccountIdTo:          accountIdTo,
		BalanceAccountIdFrom: balanceAccountIdFrom,
		BalanceAccountIdTo:   balanceAccountIdTo,
		CreatedAt:            time.Now(),
		UpdateAt:             time.Now(),
	}
	err := balance.Validate()
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (b *Balance) Validate() error {
	if b.AccountIdFrom == "" {
		return errors.New("AccountIdFrom is required")
	}
	if b.AccountIdTo == "" {
		return errors.New("AccountIdTo is required")
	}
	return nil
}
