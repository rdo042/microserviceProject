package database

import (
	"database/sql"
	"fmt"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {

	return &ClientDB{
		DB: db,
	}
}

func (c *ClientDB) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
	stmt, err := c.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}

	return client, nil

}

func (c *ClientDB) Save(client *entity.Client) error {
	fmt.Println("W1")
	stmt, err := c.DB.Prepare("INSERT INTO clients(id, name, email, created_at) VALUES (?, ?, ?, ?)")
	fmt.Println(err)
	if err != nil {
		return err
	}

	defer stmt.Close()
	fmt.Println("W3")
	_, err = stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil

}
