package models

import "fmt"

type User struct {
	ID       int64
	Name     string
	Type     string
	CPF_CNPJ string
	Email    string
	Password string
	Balance  float64
}

func ConfigQueryCreate(user User) string {
	user.Type = fmt.Sprint("'" + user.Type + "'")

	query := fmt.Sprintf("insert into users (name_user, type_user, CPF_CNPJ, email, password_user, balance values (%v, %v, %v, %v, %v, %v);", user.Name, user.Type, user.CPF_CNPJ, user.Email, user.Password, user.Balance)
	return query
}

func ConfigQueryGet(ID int64) string {
	query := fmt.Sprintf("select * from users where id = %v", ID)
	return query
}

func ConfigQueryUpdate(user User) string {
	query := fmt.Sprintf("update users set wallet = %v where id = %v ;", user.Balance, user.ID)
	return query
}
