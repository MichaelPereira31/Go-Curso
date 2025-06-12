package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	http.HandleFunc("/busca-cep", BuscaCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca CEP"))
}

func BuscaCep(cep string) (*ViaCep, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

	if error != nil {
		return nil, error
	}

	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)

	if error != nil {
		return nil, error
	}

	var data ViaCep
	error = json.Unmarshal(body, &data)

	if error != nil {
		return nil, error
	}

	return &data, nil

}