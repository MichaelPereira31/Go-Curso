package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task %d - %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	// Thread 2
	go task("A")
	// Thread 3
	go task("B")

	// Thread 4 função anonima
	go func ()  {
		for i := 0; i < 5; i++ {
			fmt.Printf("Task %d - %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	// se não tiver nada o Go vai finalizar o programa antes das goroutines terminarem

	// Aguarda o suficiente para as goroutines terminarem
	time.Sleep(12 * time.Second)
}
