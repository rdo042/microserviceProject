package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com.br/rdo042/go-app-consumer/internal/database"
	"github.com.br/rdo042/go-app-consumer/internal/entity"
	"github.com.br/rdo042/go-app-consumer/internal/usecase/create_balance"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true", "root", "root", "mysql-consumer", "3306", "walletconsumer"))

	if err != nil {
		fmt.Println("Error de conexao", err.Error())
		panic(err)
	}
	defer db.Close()

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	m.Up()

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
		"client.id":         "goapp-consumer",
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	topics := []string{"balances"}
	c.SubscribeTopics(topics, nil)

	var f *entity.KafkaConsumer

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {

			balanceDb := database.NewBalanceDB(db)
			uc := create_balance.NewCreateBalanceUseCase(balanceDb)

			err = json.Unmarshal(msg.Value, &f)

			if err != nil {
				fmt.Println("Erro ao converter em objeto " + err.Error())
			}

			output, err := uc.Execute(create_balance.CreateBalanceInputDTO{
				AccountIdFrom:        f.Payload.AccountIdFrom,
				AccountIdTo:          f.Payload.AccountIdTo,
				BalanceAccountIdFrom: f.Payload.BalanceAccountIdFrom,
				BalanceAccountIdTo:   f.Payload.BalanceAccountIdTo,
			})

			if err == nil {
				fmt.Println("Registro Gravado na base de dados " + output.AccountIdFrom)
			} else {
				fmt.Println("error consumer ao gravar", err.Error())
			}

			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
