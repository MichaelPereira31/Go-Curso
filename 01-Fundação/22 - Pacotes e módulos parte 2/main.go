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
}
