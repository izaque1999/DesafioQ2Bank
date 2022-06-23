package models

import (
	"fmt"
	"q2bank/db"
)

func UpdateUser(usuario Usuario) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	query := fmt.Sprintf("UPDATE usuario SET     SALDO = %v WHERE ID = %v ;", usuario.Saldo, usuario.ID)
	res, err := conn.Exec(query)
	if err != nil {
		fmt.Println("Erro no update", res)
		return err
	}

	return nil
}
