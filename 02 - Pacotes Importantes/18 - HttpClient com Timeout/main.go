package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// Cria um cliente HTTP com timeout de 1 segundo
	// Isso significa que se a requisição demorar mais de 1 segundo, ela será cancelada e retornará um erro de timeout.
	c := http.Client{Timeout: time.Second}
	resp, err := c.Get("http://www.google.com")
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