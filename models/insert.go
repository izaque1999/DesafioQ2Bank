package models

import (
	"fmt"
	"q2bank/db"
)

func InsertTransaction(transacao Transacao) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	query := fmt.Sprintf("INSERT INTO transacao (IDPayer, IDPayee, Valor) VALUES (%v, %v, %v,)", transacao.IDPayer, transacao.IDPayee, transacao.Valor)

	res, err := conn.Exec(query)
	if err != nil {
		fmt.Println("Erro no insert", res)
		return err
	}

	return nil
}
