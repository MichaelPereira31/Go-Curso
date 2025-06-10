package main

import "fmt"

func main() {
	salarios := map[string]int{"joao": 1000, "maria": 2000, "pedro": 3000}
	fmt.Println(salarios["joao"]) // Acessando um valor do map

	delete(salarios, "maria") // Removendo um valor do map
	fmt.Println(salarios["maria"]) // Acessando um valor que não existe no map

	for nome, salario := range salarios {
		fmt.Printf("%s: %d\n", nome, salario)
	}

	sal := make(map[string]int) // Criando um map vazio
	sal["joao"] = 1000
	sal1 := map[string]int{} // Criando um map vazio
	sal1["joao1"] = 1100

	for nome, salario := range sal {
		fmt.Printf("%s: %d\n", nome, salario)
	}

	for nome, salario := range sal1 {
		fmt.Printf("%s: %d\n", nome, salario)
	}
}