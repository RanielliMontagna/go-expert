package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

const (
	SERVER_ENDPOINT = "http://localhost:8080/quotation"
	TIMEOUT_API     = 300 * time.Millisecond
	QUOTE_FILE      = "quotation.txt"
)

type QuoteResponse struct {
	Bid float64
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_API)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", SERVER_ENDPOINT, nil)
	if err != nil {
		log.Fatalf("Error creating request: %s", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error fetching USD to BRL quotation: %s", err)
		return
	}

	defer res.Body.Close()

	read_response, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading USD to BRL quotation: %s", err)
		return
	}

	var quote QuoteResponse
	err = json.Unmarshal(read_response, &quote)
	if err != nil {
		log.Fatalf("Error parsing USD to BRL quotation: %s", err)
		return
	}

	file, err := os.Create(QUOTE_FILE)
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
		return
	}

	template, err := template.New("output").Parse("Dolar: {{.Bid}}")
	if err != nil {
		log.Fatalf("Error creating template: %s", err)
		return
	}

	file_content := QuoteResponse{Bid: quote.Bid}

	err = template.Execute(file, file_content)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
		return
	}

	log.Println("USD to BRL quotation saved to file")
	log.Printf("Dolar: %f", quote.Bid)

}
