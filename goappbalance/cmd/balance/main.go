package main

import (
	"database/sql"
	"fmt"

	"github.com.br/rdo042/goappbalance/internal/database"
	"github.com.br/rdo042/goappbalance/internal/usecase/find_balance"
	"github.com.br/rdo042/goappbalance/internal/web"
	"github.com.br/rdo042/goappbalance/internal/web/webserver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-consumer", "3306", "walletconsumer"))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	balanceDb := database.NewBalanceDB(db)
	findBalanceUseCase := find_balance.NewFindBalanceUseCase(balanceDb)
	webserver := webserver.NewWebServer(":3003")
	balanceHandler := web.NewWebBalanceHandler(*findBalanceUseCase)

	webserver.AddHandler("/balances", balanceHandler.FindBalance)

	fmt.Println("Server Balance is running in PORT 3003")

	webserver.Start()

}
