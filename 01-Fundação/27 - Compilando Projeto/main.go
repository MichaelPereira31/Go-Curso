package main

func main() {
	// IF usado para executar uma instrução somente se uma condição for verdadeira
	// No GO não existe  IF-ELSE apenas se usar IF e ELSE.
	a := 10
	b := 20
	if a > b {
		println("a é maior que b")
	} else {
		println("a é menor que b")
	}

	switch a {
		case 10:
			println("a é igual a 10, usando switch")
		case 20:
			println("a é igual a 20, usando switch")
		default:
			println("a é diferente de 10 e de 20, usando switch")
	}
	
}