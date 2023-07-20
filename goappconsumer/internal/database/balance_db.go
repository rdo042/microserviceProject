package database

import (
	"database/sql"
	"fmt"

	"github.com.br/rdo042/go-app-consumer/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {

	return &BalanceDB{
		DB: db,
	}
}

func (b *BalanceDB) Save(balance *entity.Balance) error {

	stmt, err := b.DB.Prepare("INSERT INTO balance (id, account_id_from, account_id_to, balance_account_id_from, balance_account_id_to, created_at) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Println("Error no prepare", err.Error())
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(balance.ID, balance.AccountIdFrom, balance.AccountIdTo, balance.BalanceAccountIdFrom, balance.BalanceAccountIdTo, balance.CreatedAt)

	if err != nil {
		fmt.Println("Error no gravar", err.Error())
		return err
	}

	return nil

}
