package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	API_PORT      = 8080
	TIMEOUT_API   = 250 * time.Millisecond
	TIMEOUT_DB    = 10 * time.Millisecond
	QUOTE_API_URL = "https://economia.awesomeapi.com.br/json/last"
	SOURCE_COIN   = "USD"
	TARGET_COIN   = "BRL"
)

type UsdBrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Quote struct {
	Usdbrl UsdBrl `json:"USDBRL"`
}

type QuoteResponse struct {
	Bid float64 `json:"bid" gorm:"primaryKey"`
}

type QuoteDB struct {
	ID   uint      `gorm:"primaryKey"`
	Bid  float64   `json:"bid"`
	Date time.Time `json:"date" gorm:"autoCreateTime"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/quotation", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", API_PORT), mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received on %s", r.URL.Path)

	ctx_api, cancel := context.WithTimeout(context.Background(), TIMEOUT_API)
	defer cancel()

	log.Println("Creating request")
	req, err := http.NewRequestWithContext(ctx_api, "GET", fmt.Sprintf("%s/%s-%s", QUOTE_API_URL, SOURCE_COIN, TARGET_COIN), nil)
	if err != nil {
		log.Fatalf("Error creating request: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Println("Requesting USD to BRL quotation")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error fetching USD to BRL quotation: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	log.Println("Reading USD to BRL quotation")
	read_response, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading USD to BRL quotation: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Println("Parsing USD to BRL quotation")
	var quote Quote
	err = json.Unmarshal(read_response, &quote)
	if err != nil {
		log.Fatalf("Error parsing USD to BRL quotation: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Println("Converting USD to BRL quotation to float")
	quote_float, err := strconv.ParseFloat(quote.Usdbrl.Bid, 64)
	if err != nil {
		log.Fatalf("Error converting USD to BRL quotation to float: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Println("Connecting to database")
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
		return
	}

	log.Println("Creating table if not exists")
	db.AutoMigrate(&QuoteDB{})

	ctx_db, cancel := context.WithTimeout(context.Background(), TIMEOUT_DB)
	defer cancel()

	// Insert the USD to BRL quotation in the database using the ORM and context with timeout of 10 milliseconds

	select {
	case <-ctx_db.Done():
		log.Fatalf("Error inserting USD to BRL quotation in the database: %s", ctx_db.Err())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	default:
		log.Println("Inserting USD to BRL quotation in the database")
		db.Create(&QuoteDB{Bid: quote_float})

		log.Println("Reading USD to BRL quotation from the database")
		var quote_db QuoteDB
		db.Last(&quote_db)

		log.Println("Returning USD to BRL quotation")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(QuoteResponse{Bid: quote_db.Bid})

		log.Printf("Request processed on %s", r.URL.Path)
	}
}
