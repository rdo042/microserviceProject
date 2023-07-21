package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com.br/rdo042/goappbalance/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (b *BalanceDB) Get(id string) (*entity.Balance, error) {

	var balance entity.Balance

	stmt, err := b.DB.Prepare("SELECT id, account_id_from, account_id_to, balance_account_id_from, balance_account_id_to, created_at, update_at FROM balance WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var dateUpdString sql.NullString

	row := stmt.QueryRow(id)
	err = row.Scan(
		&balance.ID,
		&balance.AccountIdFrom,
		&balance.AccountIdTo,
		&balance.BalanceAccountIdFrom,
		&balance.BalanceAccountIdTo,
		&balance.CreatedAt,
		&dateUpdString,
	)

	if dateUpdString.String != "" && dateUpdString.String != "null" {
		dateUpd, error := time.Parse("2006-01-02T15:04:05Z", dateUpdString.String)
		if error != nil {
			fmt.Println(error)

		} else {
			balance.UpdateAt = dateUpd
			fmt.Println("Data ", dateUpd)
		}
	}

	if err != nil {
		return nil, err
	}
	return &balance, nil

}
