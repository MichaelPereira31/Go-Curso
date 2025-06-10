package main

import "fmt"

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
}

func main() {
	michael := Cliente{
		Nome: "Michael",
		Idade: 25,
		Ativo: true,
	}

	michael.Ativo = false // Atualizando o status de ativo

	fmt.Printf("Nome: %s,\nIdade: %d,\nAtivo: %t\n", michael.Nome, michael.Idade, michael.Ativo)
}