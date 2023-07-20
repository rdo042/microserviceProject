package gateway

import "github.com.br/rdo042/go-app-consumer/internal/entity"

type BalanceGateway interface {
	Save(balance *entity.Balance) error
}
