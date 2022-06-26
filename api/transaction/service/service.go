package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"q2bank/api/transaction/models"
	"q2bank/api/transaction/repository"
	userModels "q2bank/api/user/models"
	userRepository "q2bank/api/user/repository"
	"strconv"
)

type URL struct {
	Authorization bool `json:"authorization"`
}

func GetTransactionID(ID userModels.User) {

}

var (
	payer               userModels.User
	payee               userModels.User
	err                 error
	mock                bool
	IDConv              int64
	transactionResponse []models.Transaction
	result              []byte
)

func MakeTransaction(MTransaction models.Transaction) (string, error) {

	mock, err = GetMock()
	if err != nil {
		return "Error in Mock", err
	}

	if mock != true {
		return "Mock is false", err
	}

	token, cond := CheckPayer(payer, MTransaction.Values)
	if token != true {

		checkMSG := fmt.Sprintf("Permission denied: %v", cond)
		return checkMSG, err
	}
	err = repository.InsertTransaction(MTransaction)
	if err != nil {
		return "Error entering transaction in Data Base", err
	}

	payer, err = userRepository.GetUserId(MTransaction.IDPayer)
	if err != nil {
		return "Payer does not exist", err
	}

	payee, err = userRepository.GetUserId(MTransaction.IDPayee)
	if err != nil {
		return "Payee does not exist", err
	}

	payer.Balance -= MTransaction.Values
	payee.Balance += MTransaction.Values

	err = userRepository.UpdateBalance(payer)
	if err != nil {
		return "Error updating values Payer", err
	}

	err = userRepository.UpdateBalance(payee)
	if err != nil {
		return "Error updating valyes Payee", err
	}

	return "Successful transaction", nil
}

func GetMock() (bool, error) {

	var url URL

	result, err := http.Get("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")
	if err != nil {
		return false, err
	}

	res, _ := ioutil.ReadAll(result.Body)

	err = json.Unmarshal(res, &url)
	if err != nil {
		return url.Authorization, err
	}

	return url.Authorization, nil
}

func GetTransaction(ID string) (string, error) {

	IDConv, err := strconv.ParseInt(ID, 0, 64)
	if err != nil {
		return "Parse error", err
	}

	transactionResponse, err = repository.GetTransactions(IDConv)
	if err != nil {
		return "Error in fetching data", err
	}

	result, err = json.Marshal(transactionResponse)
	if err != nil {
		return "Error in the marshal", err
	}
	return string(result), nil
}

func CheckPayer(payer userModels.User, Balance float64) (bool, string) {
	if payer.Balance < Balance {
		return false, "insufficient funds"
	}

	if payer.Type != "common" {
		return false, "Storekeeper cannot transfer"
	}

	return true, ""

}
