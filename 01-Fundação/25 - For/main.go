package main

func main() {
	// For
	for i := 0; i < 10; i++ {
		println(i)
	}

	// Range
	numeros := []string{"um", "dois", "tres", "quatro", "cinco"}
	for k, n := range numeros {
		println(k, n)
	}

	// For com condição
	i := 0
	for ; i < 10; i++ {
		println(i)
	}

	//for infinito
	for {
		println("infinito")
	}
}