package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(10, 20)

	if(err != nil){
		fmt.Println(err)
	}

	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	if(a + b >= 50){
		return 0, errors.New("soma maior que 50, error")
	}
	return a + b, nil
}