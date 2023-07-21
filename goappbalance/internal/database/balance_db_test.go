package database

import (
	"database/sql"
	"fmt"
	"testing"

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

func (s *BalanceDBTestSuite) TearDownSuite() {
	defer s.db.Close()
}

func TestBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}

func (s *BalanceDBTestSuite) TestFindByID() {
	balanceDB, err := s.balanceDB.Get("0276dc9e-1d30-41e9-8225-f2d251655f05")
	fmt.Println(err)
	s.Nil(err)
	s.Equal(balanceDB.ID, "0276dc9e-1d30-41e9-8225-f2d251655f05")
	s.Equal(balanceDB.AccountIdFrom, "4b9633e7-21bf-413c-be9c-4cded7468652")
	s.Equal(balanceDB.AccountIdTo, "c990b127-9c96-4ea9-a996-c70c3d9bd44c")
}
