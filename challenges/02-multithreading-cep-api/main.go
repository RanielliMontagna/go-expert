package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
)

const (
	API_PORT    = 8080
	TIMEOUT_API = 1 * time.Second
	VIA_CEP_URL = "https://viacep.com.br/ws/%s/json/"
	BRASIL_API  = "https://brasilapi.com.br/api/cep/v1/%s"
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

type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type Cep struct {
	Cep        string `json:"cep"`
	Cidade     string `json:"cidade"`
	Uf         string `json:"uf"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Servico    string `json:"servico"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cep/", handler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", API_PORT), mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Path[len("/cep/"):]

	if cep == "" {
		http.Error(w, "cep is required", http.StatusBadRequest)
		return
	}

	if !regexp.MustCompile(`^\d{8}$`).MatchString(cep) {
		http.Error(w, "cep is invalid", http.StatusBadRequest)
		return
	}

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		viaCep, err := getViaCep(cep)
		if err != nil {
			c1 <- err.Error()
			return
		}
		c1 <- fmt.Sprintf("%+v", viaCep)
	}()

	go func() {
		brasilApi, err := getBrasilApi(cep)
		if err != nil {
			c2 <- err.Error()
			return
		}
		c2 <- fmt.Sprintf("%+v", brasilApi)
	}()

	select {
	case viaCep := <-c1:
		fmt.Fprintln(w, viaCep)
	case brasilApi := <-c2:
		fmt.Fprintln(w, brasilApi)
	case <-time.After(TIMEOUT_API):
		http.Error(w, "timeout", http.StatusRequestTimeout)
	}
}

func getViaCep(cep string) (Cep, error) {
	req, err := http.Get(fmt.Sprintf(VIA_CEP_URL, cep))
	if err != nil {
		return Cep{}, err
	}
	defer req.Body.Close()

	var data ViaCEP

	err = json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return Cep{}, err
	}

	return Cep{
		Cep:        data.Cep,
		Cidade:     data.Localidade,
		Uf:         data.Uf,
		Logradouro: data.Logradouro,
		Bairro:     data.Bairro,
		Servico:    "ViaCEP",
	}, nil
}

func getBrasilApi(cep string) (Cep, error) {
	req, err := http.Get(fmt.Sprintf(BRASIL_API, cep))
	if err != nil {
		return Cep{}, err
	}
	defer req.Body.Close()

	var data BrasilAPI
	err = json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return Cep{}, err
	}

	return Cep{
		Cep:        data.Cep,
		Cidade:     data.City,
		Uf:         data.State,
		Logradouro: data.Street,
		Bairro:     data.Neighborhood,
		Servico:    "BrasilAPI",
	}, nil
}
