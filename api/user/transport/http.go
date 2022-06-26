package transport

import (
	"fmt"
	"net/http"
	"q2bank/api/user/models"
	"q2bank/api/user/service"
	"strconv"
)

var (
	demand          models.User
	ID, result, msg string
	err             error
	msgErro         string
)

func GetUserIDHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		fmt.Fprint(w, "Method is not GET")
		return
	}

	ID = r.FormValue("ID")

	result, err = service.GetUserID(ID)
	if err != nil {
		fmt.Print("Erro in getuser", err)
		return
	}
	fmt.Fprintf(w, "%v", result)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	msgErro = "Value not null"

	if r.Method != http.MethodPost {
		fmt.Fprint(w, "Method is not POST")
		return
	}

	if demand.Name = r.FormValue("name"); demand.Name == "" {
		return
	}

	if demand.Type = r.FormValue("type"); demand.Type == "" {
		return
	}

	if demand.CPF_CNPJ = r.FormValue("cpf_cnpj"); demand.CPF_CNPJ == "" {
		return
	}

	if demand.Email = r.FormValue("email"); demand.Email == "" {
		return
	}

	if demand.Password = r.FormValue("password"); demand.Password == "" {
		return
	}

	demand.Balance, err = strconv.ParseFloat(r.FormValue("balance"), 64)
	if err != nil {
		fmt.Fprint(w, "Erro in conversion")
		return
	}

	result, err = service.CreateUser(demand)
	if err != nil {
		fmt.Fprint(w, result)
		return
	}

	fmt.Fprint(w, result)

}
