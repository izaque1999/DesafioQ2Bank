package models

import (
	"fmt"
	"time"
)

type Transaction struct {
	IDPayer int64   `json:"idpayer"`
	IDPayee int64   `json:"idpayee"`
	Values  float64 `json:"valor"`
	Date    string  `json:"date"`
}

func ConfigQueryInsert(transaction Transaction) string {
	dateFormat := fmt.Sprintf("'%v'", time.Now().Format("2006-01-02 15:04:05"))
	transaction.Date = dateFormat

	query := fmt.Sprintf("insert into transactions (values, id_payer, id_payee, date_time) values (%v, %v, %v, %v)", transaction.Values, transaction.IDPayer, transaction.IDPayee, transaction.Date)
	return query
}

func ConfigQueryGet(ID int64) string {

	query := fmt.Sprintf("select id_payer, id_payee, date_time, values from transactions where id_payer = %v or id_payee = %v", ID, ID)
	return query
}
