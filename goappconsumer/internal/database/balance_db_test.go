package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	
)

type BalanceDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	balanceDB *BalanceDB
}

func (s *BalanceDBTestSuite) SetupSuite() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-consumer", "3307", "walletconsumer"))

	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE balance (
		id integer not null auto_increment,
		account_id_from varchar(100) not null,
		account_id_to varchar(100) not null,
		balance_account_id_from decimal(19,2),
		balance_account_id_to decimal(19,2),
		PRIMARY KEY (id)
	")
	s.balanceDB = NewBalanceDB(db)
}

func (s *BalanceDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE balance")
}

func TestBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}

func (s *BalanceDBTestSuite) TestSave() {
	balance := entity.NewBalance("111111", "22222", 10, 20)
	err := s.balanceDB.Save(balance)

	s.Nil(err)

}