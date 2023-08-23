package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/nectariosouzou/portfolio-viewer/backend/api"
	"github.com/nectariosouzou/portfolio-viewer/backend/infra"
	"github.com/rs/cors"
)

type RequestData struct {
	Data string `json:"data"`
}

const port = ":8080"

var dbs *infra.DbClients
var initDbOnce sync.Once

func getDbs() *infra.DbClients {
	initDbOnce.Do(func() {
		dbs = infra.InitDbClients()
	})
	return dbs
}

func configureCors() *cors.Cors {
	options := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	}
	return cors.New(options)
}

func shutdownHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			getDbs().RedisClient.Close()
			os.Exit(0)
		}
	}()
}

func main() {
	shutdownHandler()
	infra.InitEnv()
	corsMiddleware := configureCors()
	wrappedHandler := corsMiddleware.Handler(http.DefaultServeMux)

	app := api.InitApi(getDbs())
	// API routes
	http.HandleFunc("/portfolio", app.Portfolio)
	fmt.Println("Server is running on port" + port)
	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, wrappedHandler))
}
