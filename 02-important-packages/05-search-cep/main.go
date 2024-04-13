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
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", cep, err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: parsing %s: %v\n", cep, err)
		}

		file, err := os.Create("city.txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: creating file: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Logradouro: %s, Complemento: %s, Bairro: %s, Localidade: %s, UF: %s, IBGE: %s, GIA: %s, DDD: %s, SIAFI: %s\n", data.Cep, data.Logradouro, data.Complemento, data.Bairro, data.Localidade, data.Uf, data.Ibge, data.Gia, data.Ddd, data.Siafi))

		fmt.Println("City data saved in city.txt")
		fmt.Println("CEP: ", data.Cep)
		fmt.Println("Localidade: ", data.Localidade)
		fmt.Println("UF: ", data.Uf)
	}
}
