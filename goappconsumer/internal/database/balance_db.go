package database

import (
	"database/sql"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (t *BalanceDB) Create(balance *entity.Balance) error {

	stmt, err := t.DB.Prepare("INSERT INTO balance (account_id_from, account_id_to, balance_account_id_from, balance_account_id_to) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(balance.AccountFrom.ID, balance.AccountTo.ID, balance.BalanceAccountIdFrom, balance.BalanceAccountIdTo)
	if err != nil {
		return err
	}

	return nil

}
