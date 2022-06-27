package models

import (
	"fmt"
	"time"
)

type Transaction struct {
	IDPayer int64
	IDPayee int64
	Values  float64
	Date    string
}

func ConfigQueryInsert(transaction Transaction) string {
	dateFormat := fmt.Sprintf("'%v'", time.Now().Format("2006-01-02 15:04:05"))
	transaction.Date = dateFormat

	query := fmt.Sprintf("insert into transactions (values_transaction, id_payer, id_payee, date_transaction) values (%v, %v, %v, %v)", transaction.Values, transaction.IDPayer, transaction.IDPayee, transaction.Date)
	return query
}

func ConfigQueryGet(ID int64) string {

	query := fmt.Sprintf("select id_payer, id_payee, date_transaction, values from transactions where id_payer = %v or id_payee = %v", ID, ID)
	return query
}
