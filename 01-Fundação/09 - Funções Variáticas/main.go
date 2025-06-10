package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(10, 20, 30, 40, 50))
}

func sum(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}
