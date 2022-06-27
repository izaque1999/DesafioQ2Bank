package main

import (
	"net/http"
	transactionTransport "q2bank/api/transaction/transport"
	"q2bank/api/user/transport"
)

func main() {

	//Router user
	http.HandleFunc("/getUserID", transport.GetUserIDHandler)
	http.HandleFunc("/createUser", transport.CreateUserHandler)

	//Router transaction
	http.HandleFunc("/getTransactions", transactionTransport.GetTransactionsHandler)
	http.HandleFunc("/transaction", transactionTransport.TransactionHandler)

	http.ListenAndServe(":9000", nil)
}
