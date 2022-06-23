package models

import "q2bank/db"

func Get(id int64) (usuario Usuario, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM usuario WHERE id=$1`, id)

	err = row.Scan(&usuario.ID, &usuario.Nome, &usuario.Tipo, &usuario.Cpf, &usuario.Email)

	return
}
