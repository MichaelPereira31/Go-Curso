package main

import "fmt"

func main() {
	a := "Hello, World!" // na primeira vez é atribuição
	a = "Hello, World"   // atribuição
	fmt.Printf("O tipo de a é: %T", a)
	fmt.Printf("O valor de a é: %v", a)
}
