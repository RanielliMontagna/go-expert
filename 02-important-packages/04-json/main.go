package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance"`
}

func main() {
	conta := Account{Number: 1, Balance: 100}

	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON: %s\n", res)

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	pureJson := []byte(`{"number":2,"balance":200}`)
	var accountX Account

	err = json.Unmarshal(pureJson, &accountX)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Account X: %+v\n", accountX)
}
