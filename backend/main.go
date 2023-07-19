package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/signal"

	"github.com/rs/cors"
)

var db *DbClients

func InitDb() *DbClients {
	return InitDbClients()
}

type RequestData struct {
	Data string `json:"data"`
}

func main() {
	db = InitDb()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			db.RedisClient.Close()
			os.Exit(0)
		}
	}()
	options := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
		//AllowedHeaders: []string{"Content-Type"},
	}
	corsMiddleware := cors.New(options)
	wrappedHandler := corsMiddleware.Handler(http.DefaultServeMux)
	// API routes
	http.HandleFunc("/portfolio", func(w http.ResponseWriter, r *http.Request) {
		// Parse the multipart form data
		err := r.ParseMultipartForm(32 << 20) // Max size: 32MB
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		// Get the uploaded file
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Failed to get file from form data", http.StatusBadRequest)
			return
		}
		defer file.Close()
		data := portfolio(file)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		// Encode the Message struct as JSON and write it to the response
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	port := ":5050"
	fmt.Println("Server is running on port" + port)
	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, wrappedHandler))
}

func portfolio(file multipart.File) *PortfolioValue {
	stocks, _ := processStocks(file)
	sectors, err := sortStocks(stocks)
	if err != nil {
		fmt.Printf("Error in sorting stocks: %s", err)
	}
	portfolioValues := calculatePortfolio(stocks, sectors)
	return portfolioValues
}
