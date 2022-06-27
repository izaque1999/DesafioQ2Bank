package repository

import (
	"fmt"
	"q2bank/api/db"
	"q2bank/api/user/models"

	_ "github.com/go-sql-driver/mysql"
)

var (
	query string
	err   error
)

func CreateUser(user models.User) (string, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return "Error opening database", nil
	}

	defer conn.Close()

	query = models.ConfigQueryCreate(user)
	_, err = conn.Query(query)
	if err != nil {
		return "Error executing query", err
	}

	return "successful query", err
}

func GetUserId(ID int64) (models.User, error) {
	fmt.Println("Entrando na função")
	var userID models.User

	conn, err := db.OpenConnection()

	if err != nil {
		fmt.Println("1")

		return userID, nil
	}

	defer conn.Close()

	query = models.ConfigQueryGet(ID)

	ret, err := conn.Query(query)

	if err != nil {
		fmt.Println("2")

		return userID, nil
	}

	ret.Next()
	err = ret.Scan(&userID.ID, &userID.Name, &userID.Type, &userID.CPF_CNPJ, &userID.Email, &userID.Balance, &userID.Balance)
	if err != nil {
		fmt.Println("3")

		return userID, nil
	}

	fmt.Printf("retorno: %v", ret)

	return userID, nil
}

func UpdateBalance(userUp models.User) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}

	query = models.ConfigQueryUpdate(userUp)

	_, err = conn.Exec(query)

	defer conn.Close()

	return nil
}
