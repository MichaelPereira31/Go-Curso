package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3


	for i, v:= range meuArray {
		fmt.Printf("O valor de meuArray[%d] é: %d\n", i, v)
	}

}
