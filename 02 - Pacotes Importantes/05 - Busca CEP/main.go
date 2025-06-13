package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v \n", err)
		}

		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o conteúdo: %v \n", err)
		}

		var data ViaCep
		err = json.Unmarshal(res, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar o conteúdo: %v \n", err)
		}

		fmt.Println(data)

		file, err := os.Create("05 - Busca CEP/cidade.txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo: %v \n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s \n", data.Cep, data.Localidade, data.Uf))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v \n", err)
		}

	}
}