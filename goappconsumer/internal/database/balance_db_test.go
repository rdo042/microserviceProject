package database

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com.br/rdo042/go-app-consumer/internal/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

type BalanceDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	balanceDB *BalanceDB
}

func (s *BalanceDBTestSuite) SetupSuite() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-consumer", "3306", "walletconsumer"))

	s.Nil(err)
	s.db = db
	s.balanceDB = NewBalanceDB(db)
}

func TestBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}

func (s *BalanceDBTestSuite) TestSave() {
	balance, _ := entity.NewBalance("111111", "22222", 10, 20)
	err := s.balanceDB.Save(balance)

	s.Nil(err)

}
