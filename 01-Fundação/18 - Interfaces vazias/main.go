package main

import "fmt"

func main() {
	// Interfaces vazias são usadas quando não se sabe o tipo exato de dado que será passado
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)
}

func showType(i interface{}){
	fmt.Printf("O tipo da variável é: %T, o valor da variável é: %v\n", i, i)
}