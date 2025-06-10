package main
func main() {
	// Memória -> Endereço -> Valor
	// Ponteiro é uma variável que armazena o endereço de memória de outra variável
	a := 10
	println("Valor de a:", a)
	var ponteiro *int = &a // p é um ponteiro que aponta para o endereço de a
	println("Endereço de a:", ponteiro)
	*ponteiro = 20 // Atribuindo um novo valor a a
	println("Valor de a:", a)
	b:= &a // b é um ponteiro que aponta para o endereço de a
	println("Endereço de a:", b)
}