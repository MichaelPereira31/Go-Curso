package main

import (
	"fmt"
)

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

type Pessoa interface {
	Desativar() // Método que deve ser implementado por qualquer tipo que implemente a interface Pessoa
}

func Desativacao(pessoa Pessoa){
	pessoa.Desativar() // Chamando o método Desativar
}

func (c Cliente) Desativar(){
	c.Ativo = false // Método para desativar o cliente
	fmt.Println("Cliente desativado:", c.Nome)
}

func main() {
	michael := Cliente{
		Nome: "Michael",
		Idade: 25,
		Ativo: true,
	}

	Desativacao(michael) // Chamando a função Desativacao com o objeto Cliente
}