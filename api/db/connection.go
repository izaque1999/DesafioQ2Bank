package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:12300596Izaque@tcp(localhost)/q2bank")
	if err != nil {
		fmt.Println("Erro na conexão: ", err)
		return db, err
	}
	defer db.Close()

	return db, nil
}
