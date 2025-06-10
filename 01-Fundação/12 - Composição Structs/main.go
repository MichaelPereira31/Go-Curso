package main

import "fmt"

type Endereco struct {
	Logradouro string
	Cidade     string
	Estado     string
	Numero		 int
}

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
	Endereco  Endereco// Composição: Cliente tem um Endereco

}

func main() {
	michael := Cliente{
		Nome: "Michael",
		Idade: 25,
		Ativo: true,
	}

	michael.Ativo = false // Atualizando o status de ativo
	michael.Endereco.Cidade = "São Paulo" // Acessando o campo Cidade do Endereco
	michael.Endereco.Logradouro = "Rua Exemplo"

	fmt.Printf("Nome: %s,\nIdade: %d,\nAtivo: %t\n", michael.Nome, michael.Idade, michael.Ativo)
	fmt.Printf("Cidade: %s,\nLogradouro: %s\n", michael.Endereco.Cidade, michael.Endereco.Logradouro)
}