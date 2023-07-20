package create_balance

import (
	"time"

	"github.com.br/rdo042/go-app-consumer/internal/entity"
	"github.com.br/rdo042/go-app-consumer/internal/gateway"
)

type CreateBalanceInputDTO struct {
	AccountIdFrom        string  `json:"account_id_from"`
	AccountIdTo          string  `json:"account_id_to"`
	BalanceAccountIdFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIdTo   float64 `json:"balance_account_id_to"`
}

type CreateBalanceOutputDTO struct {
	ID                   string
	AccountIdFrom        string
	AccountIdTo          string
	BalanceAccountIdFrom float64
	BalanceAccountIdTo   float64
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type CreateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewCreateBalanceUseCase(balanceGateway gateway.BalanceGateway) *CreateBalanceUseCase {
	return &CreateBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (uc *CreateBalanceUseCase) Execute(input CreateBalanceInputDTO) (*CreateBalanceOutputDTO, error) {

	balance, err := entity.NewBalance(input.AccountIdFrom, input.AccountIdTo, input.BalanceAccountIdFrom, input.BalanceAccountIdTo)
	if err != nil {
		return nil, err
	}

	err = uc.BalanceGateway.Save(balance)
	if err != nil {
		return nil, err
	}

	output := &CreateBalanceOutputDTO{
		ID:                   balance.ID,
		AccountIdFrom:        balance.AccountIdFrom,
		AccountIdTo:          balance.AccountIdTo,
		BalanceAccountIdFrom: balance.BalanceAccountIdFrom,
		BalanceAccountIdTo:   balance.BalanceAccountIdTo,
		CreatedAt:            balance.CreatedAt,
		UpdatedAt:            balance.UpdateAt,
	}

	return output, nil
}
