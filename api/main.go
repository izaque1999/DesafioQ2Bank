package main

import (
	"q2bank/api/db"
	"q2bank/api/models"
)

type url struct {
	Authorized string
}

var user models.Usuario
var transaction models.Transacao

func main() {

	transaction = models.Transacao{
		IDPayer: 2,
		IDPayee: 3,

		Valor: 50,
	}

	//err := configs.Load()
	//if err != nil {
	//	panic(err)
	//}
	//models.InsertTransaction(transaction)
	//models.UpdateUser(user)
	//router := chi.NewRouter()
	//router.Post("/", handlers.Create)
	//router.Put("/{id}", handlers.UpdateHandler)
	//router.Delete("/{id}", handlers.Delete)
	//router.Get("/", handlers.List)
	//router.Get("/", handlers.Get)

	//link, err := http.Get("")
	//	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router) */
	db.OpenConnection()
}
