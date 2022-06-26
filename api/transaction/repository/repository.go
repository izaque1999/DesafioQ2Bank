package repository

import (
	"fmt"
	"q2bank/api/db"
	"q2bank/api/transaction/models"
)

var (
	tr           models.Transaction
	transactions []models.Transaction
	query        string
)

func GetTransactions(ID int64) ([]models.Transaction, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query = models.ConfigQueryGet(ID)

	ret, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for ret.Next() {
		err = ret.Scan(&tr.IDPayer, &tr.IDPayee, &tr.Date, &tr.Values)
		if err != nil {
			fmt.Printf("Erro ao verificar dados %v\n", err)
			return nil, err
		}
		transactions = append(transactions, tr)
	}

	return transactions, nil
}

func InsertTransaction(transactionInsert models.Transaction) error {

	return nil
}
