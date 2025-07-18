package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main(){
	// Cria um contexto com timeout de 1 segundo
	// Isso significa que se a requisição demorar mais de 1 segundo, ela será cancelada e retornará um erro de timeout.
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))


}