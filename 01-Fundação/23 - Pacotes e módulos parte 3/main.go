package main

// Inicializando o módulo
// go mod init github.com/MichaelPereira31/go-curso

// Instalando o módulo
// go mod tidy

import (
	"fmt"
	"github.com/MichaelPereira31/go-curso/matematica"
)

func main() {

	soma := matematica.Soma(10, 20)
	fmt.Println("Resultado da soma: ", soma)

	// O valor de a é acessível apenas dentro do módulo matematica, não dentro do módulo main
	// fmt.Println("Valor de a:", matematica.a)

	// Para acessar o valor de a, é necessário renomeá-lo para A assim o valor ficara acessível em todo o módulo
	fmt.Println("Valor de a:", matematica.A)
}
