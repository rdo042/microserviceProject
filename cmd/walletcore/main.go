package main

import (
	"database/sql"
	"fmt"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/database"
	//"github.com.br/devfullcycle/fc-ms-wallet/internal/event"
	//"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_account"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_client"
	//"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web/webserver"

	//"github.com.br/devfullcycle/fc-ms-wallet/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//eventDispatcher := events.NewEventDispatcher()
	//transactionCreatedEvent := event.NewTransactionCreated()

	//eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)

	//accountDb := database.NewAccountDB(db)
	//transactionDb := database.NewTransactionDB(db)

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)

	//createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	//createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	//accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	//transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	//webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	//webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")

	webserver.Start()
}