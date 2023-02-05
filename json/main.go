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
	acc := Account{Number: 1, Balance: 1000}

	res, err := json.Marshal(acc)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(acc)
	if err != nil {
		println(err)
	}

	rawJson := []byte(`{"number":2, "balance": 1000}`)
	var account Account
	err = json.Unmarshal(rawJson, &account)
	if err != nil {
		panic(err)
	}
	fmt.Println(account)
}
