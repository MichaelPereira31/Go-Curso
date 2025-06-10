package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Cria um arquivo chamado arquivo.txt
	f, error := os.Create("01 - Manipulação de Arquivos/arquivo.txt")

	if error != nil {
		panic(error)
	}

	// Escreve no arquivo
	// tamanho, error := f.WriteString("Olá mundo!")
	tamanho, error := f.Write([]byte("Escrevendo no arquivo"))

	if error != nil {
		panic(error)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", tamanho)
	f.Close()

	// Abre o arquivo e lê o conteúdo
	arquivo, error := os.ReadFile("01 - Manipulação de Arquivos/arquivo.txt")

	if error != nil {
		panic(error)
	}

	fmt.Printf("Conteúdo do arquivo: %s\n", string(arquivo))

	// Leitura de pouco em pouco o conteúdo do arquivo
	arquivo2, error := os.Open("01 - Manipulação de Arquivos/arquivo.txt")

	if error != nil {
		panic(error)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)

	for {
		n, error := reader.Read(buffer)

		if error != nil{
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	arquivo2.Close()

	// Deleta o arquivo
	// error = os.Remove("01 - Manipulação de Arquivos/arquivo.txt")

	// if error != nil {
	// 	panic(error)
	// }

}