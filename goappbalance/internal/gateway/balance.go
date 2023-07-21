package gateway

import "github.com.br/rdo042/goappbalance/internal/entity"

type BalanceGateway interface {
	Get(id string) (*entity.Balance, error)
}
