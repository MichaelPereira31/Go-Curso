package main

import "fmt"

func main() {
	// Utilizando defer para executar quando finalizar a função
	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}