package main

import "fmt"


func main() {
	// Type assertation é usado para verificar se um valor é do tipo esperado
	var minhaVar interface{} = "Olá, Mundo!"

	// Verifica se minhaVar é do tipo string
	println("Valor original:", minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Printf("Valor: %v, Tipo: %T, Verificação: %t\n", res, res, ok)
}
