# microserviceProject
Microservices Project in Go

## goapp
    Container contendo a aplicação GO responsável por realizar os cadastros de Clientes, contas e transações.
    Ao realizar uma transação, ocorre a comunicação com o Kafka.
    Para testar pode ser utilizado o arquivo /goapp/api/client.http
    De ser realizado o start da aplicacao.
    Porta: 8080

    # docker compose exec goapp bash
    # go run cmd/walletcore/main.go

## goappconsumer
    Container contendo a aplicação GO responsável por realizar a gravação dos retornos dos balances no banco de dados mysql-consumer.
    Recebe o retorno do kafka através do consumer e grava na base de dados
    De ser realizado o start da aplicacao.
    Porta: 3005

    # docker compose exec goappconsumer bash
    # go run cmd/consumer/main.go

## goappbalance
    Container contendo a aplicação GO responsável por retornar os dados dos balances.
    Obtem o resultado de acordo com a base de dados mysqlconsumer
    Porta: 3003

    # docker compose exec goappbalance bash
    # go run cmd/balance/main.go

    Para testar pode ser utilizado o arquivo /goappbalance/api/client.http

## mysql
    Contém o banco de dados wallet, que mantém as informações da aplicação goapp.
    Porta: 3306

    # docker compose exec mysql bash
    # mysql -u root -p wallet
    # **** (usar senha root)
    # show tables;

## mysql-consumer:
    Contém o banco de dados walletconsumer, que mantém as informações retornadas pelo kafka.
    Porta: 33040

    # docker compose exec mysql-consumer bash
    # mysql -u root -p walletconsumer
    # **** (usar senha root)
    # show tables;
    # select * from balance;


## zookeeper
    Container para utilização da Lib confluence para o Kafka
## kafka
    Container kafka
## control-center
    Abrir no browser http://localhost:9021



