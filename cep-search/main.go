package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json", cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when requesting: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when reading response: %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when parsing response: %v\n", err)
		}

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when creating file: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, City: %s", data.Cep, data.Localidade))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when writing in the file: %v\n", err)
		}
	}
}
