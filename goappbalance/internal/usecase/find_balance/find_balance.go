package find_balance

import (
	"time"

	"github.com.br/rdo042/goappbalance/internal/gateway"
)

type FindBalanceInputDTO struct {
	BalanceID string `json:"balance_id"`
}

type FindBalanceOutputDTO struct {
	ID                   string
	AccountIdFrom        string
	AccountIdTo          string
	BalanceAccountIdFrom float64
	BalanceAccountIdTo   float64
	CreatedAt            time.Time
}

type FindBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewFindBalanceUseCase(b gateway.BalanceGateway) *FindBalanceUseCase {
	return &FindBalanceUseCase{
		BalanceGateway: b,
	}
}

func (uc *FindBalanceUseCase) Execute(input FindBalanceInputDTO) (*FindBalanceOutputDTO, error) {
	balance, err := uc.BalanceGateway.Get(input.BalanceID)
	if err != nil {
		return nil, err
	}

	output := &FindBalanceOutputDTO{
		ID:                   balance.ID,
		AccountIdFrom:        balance.AccountIdFrom,
		AccountIdTo:          balance.AccountIdTo,
		BalanceAccountIdFrom: balance.BalanceAccountIdFrom,
		BalanceAccountIdTo:   balance.BalanceAccountIdTo,
		CreatedAt:            balance.CreatedAt,
	}
	return output, nil
}
