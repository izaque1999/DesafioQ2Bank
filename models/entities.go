package models

type Usuario struct {
	ID    int64   `json:"id"`
	Nome  string  `json:"nome"`
	Tipo  string  `json:"tipo"`
	Cpf   string  `json:"cpf"`
	Email string  `json:"email"`
	Saldo float64 `json:"saldo"`
}

type Transacao struct {
	ID      int64   `json:"id"`
	IDPayer int64   `json:"idpayer"`
	IDPayee int64   `json:"idpayee"`
	Valor   float64 `json:"valor"`
}
