package main

//Thread 1
func main(){
	// utlizada para que uma thread se comunique com a outra, ou seja, para que uma thread envie uma mensagem para a outra, ou para que uma thread espere uma mensagem da outra
	forever := make(chan bool)

	go func(){
		for i := 0; i < 10; i++ {
			println(i)
		}
		//Sempre dever ser rodado em um outra goroutine, para não bloquear a execução do programa
		forever <- true
	}()

	<- forever
}