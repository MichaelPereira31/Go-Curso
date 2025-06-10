package main

// A função soma recebe dois ponteiros para inteiros e retorna a soma dos valores apontados pelos ponteiros.
func soma(a, b *int) int {
	// Atribuindo um novo valor ao primeiro ponteiro
	// Isso altera o valor da variável original que o ponteiro aponta.
	*a = 50
	*b = 60
	return *a + *b
}
// A função soma recebe dois ponteiros para inteiros, altera o valor do primeiro ponteiro e retorna a soma dos valores apontados pelos ponteiros.
func main(){
	minhaVar1 := 10
	minhaVar2 := 20
	println("Valor da soma:", soma(&minhaVar1, &minhaVar2))

	println("Valor de minhaVar1:", minhaVar1) // Exibe o novo valor de minhaVar1 após a chamada da função soma e alteração do valor do ponteiro
	println("Valor de minhaVar2:", minhaVar2) // Exibe o novo valor de minhaVar2 após a chamada da função soma e alteração do valor do ponteiro
}