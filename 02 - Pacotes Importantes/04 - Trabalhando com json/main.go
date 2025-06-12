package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	// Retorno do json é retornado em bytes
	res, err := json.Marshal(conta)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))

	// Decodificação de bytes em um objeto e escreve no console
	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		fmt.Println(err)
	}

	jsonPuro := []byte(`{"numero":1,"saldo":100}`)
	var contaX Conta

	// usando o unmarshal para decodificar o conteúdo do json em um objeto Conta e escreve no console
	err = json.Unmarshal(jsonPuro, &contaX)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(contaX.Saldo)
}
