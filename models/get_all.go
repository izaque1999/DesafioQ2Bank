package models

import "q2bank/db"

func GetAll() (usuarios []Usuario, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM usuario`)
	if err != nil {
		return
	}
	for rows.Next() {
		var usuario Usuario

		err = rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Tipo, &usuario.Cpf, &usuario.Email)
		if err != nil {
			continue
		}

		usuarios = append(usuarios, usuario)
	}
	return
}
