package transport

import (
	"fmt"
	"log"
	"net/http"
	"q2bank/api/transaction/models"
	"q2bank/api/transaction/service"
	"strconv"
)

var (
	demand          models.Transaction
	err             error
	msg, result, ID string
	msgErro         int
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Fprint(w, "Method is not POST")
		return
	}

	demand.IDPayer, err = strconv.ParseInt(r.FormValue("payer"), 0, 64)
	if err != nil {
		fmt.Fprint(w, "Error to convert payer id")
		return
	}

	demand.IDPayee, err = strconv.ParseInt(r.FormValue("payee"), 0, 64)
	if err != nil {
		fmt.Fprint(w, "Error to convert payee id")
		return
	}

	if demand.IDPayer == demand.IDPayer {
		fmt.Fprint(w, "Impossible to transfer to yourself")
		return
	}
	demand.Values, err = strconv.ParseFloat(r.FormValue("values"), 64)
	if err != nil {
		log.Print("Error to convert Values")
		return
	}

	msg, err = service.MakeTransaction(demand)

	fmt.Fprintf(w, "%v", msg)

}

func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		fmt.Fprint(w, "Method is not GET")
		return
	}

	ID = r.FormValue("ID")

	result, err = service.GetTransaction(ID)

	if err != nil {
		fmt.Print(err)
		return
	}
	if result == "" {
		fmt.Print(err)
		return
	}

	fmt.Fprint(w, result)
}
