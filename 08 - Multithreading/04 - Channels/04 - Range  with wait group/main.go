package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go pusblish(ch)
	go reader(ch, &wg)
	wg.Wait()

}
func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}

func pusblish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
