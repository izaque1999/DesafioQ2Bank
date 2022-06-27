package service

import (
	"encoding/json"
	"fmt"
	"q2bank/api/user/models"
	"q2bank/api/user/repository"
	"strconv"
)

var (
	IDConv     int64
	err        error
	resultB    []byte
	resultS    string
	userResult models.User
)

func CreateUser(userCreate models.User) (string, error) {
	fmt.Printf("UserCreate: %v ", userCreate)
	switch {
	case userCreate.Type == "storekeeper":
		break
	case userCreate.Type == "common":
		break
	default:
		return "Invalid Type", err
	}

	resultS, err = repository.CreateUser(userCreate)
	if err != nil {
		return resultS, err
	}

	return "User registered sucessfully", err

}

func GetUserID(ID string) (string, error) {
	if ID == "" {
		return "empty string", err
	}

	IDConv, err = strconv.ParseInt(ID, 0, 64)
	if err != nil {
		return "Parse error", err
	}

	userResult, err = repository.GetUserId(IDConv)
	if err != nil {
		return "GetUserID error", err
	}

	userResult.Password = ""
	resultB, err = json.Marshal(userResult)

	return string(resultB), nil
}
