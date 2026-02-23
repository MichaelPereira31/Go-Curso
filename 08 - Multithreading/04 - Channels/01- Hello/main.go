package main

// Thread 1
func main(){
	canal := make(chan string)

	// Thread 2
	go func() {
		canal <- "Hello World"
	}()

	// Thread 3
	mensagem := <- canal
	println(mensagem)
}