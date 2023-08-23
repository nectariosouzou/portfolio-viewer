package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nectariosouzou/portfolio-viewer/backend/handlers"
	"github.com/nectariosouzou/portfolio-viewer/backend/infra"
)

type Api struct {
	RedisHandler handlers.RedisHandler
	GptHandler   handlers.GptHandler
}

type SectorInfo struct {
	Sector string
	Stocks []Stock
	Value  float32
}

type PortfolioValue struct {
	TotalValue       float32
	SectorPercentage map[string]float32
	Sectors          []SectorInfo
}

func InitApi(dbs *infra.DbClients) *Api {
	return &Api{
		RedisHandler: handlers.InitRedisHandler(dbs.RedisClient),
	}
}

func (a *Api) Portfolio(w http.ResponseWriter, r *http.Request) {
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
	stocks, _ := processStocks(file)
	sectors, err := a.sortStocks(stocks)
	if err != nil {
		fmt.Printf("Error in sorting stocks: %s", err)
	}
	portfolioValues := a.calculatePortfolio(stocks, sectors)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// Encode the Message struct as JSON and write it to the response
	err = json.NewEncoder(w).Encode(portfolioValues)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Api) sortStocks(stocks []Stock) (map[string]string, error) {
	tickers := make(map[string]bool, len(stocks))
	var tickersForGPT []string
	var sectors map[string]string
	for _, stock := range stocks {
		tickers[stock.Ticker] = false
	}
	sectors, tickers, err := a.RedisHandler.FindSectors(tickers)
	if err != nil {
		return nil, err
	}
	for key, ticker := range tickers {
		if !ticker {
			tickersForGPT = append(tickersForGPT, key)
		}
	}
	if len(tickersForGPT) > 0 {
		sectorsGPT, err := a.GptHandler.FindSectors(tickersForGPT)
		if err != nil {
			return nil, err
		}
		err = a.RedisHandler.SetTicker(sectorsGPT)
		if err != nil {
			//Don't fail. Report, new tickers won't be added to cache. TODO: Maybe log to file.
			log.Print(err)
		}
		for key, sector := range sectorsGPT {
			sectors[key] = sector
		}
	}
	return sectors, nil
}

func (a *Api) calculatePortfolio(stocks []Stock, sectors map[string]string) *PortfolioValue {
	sectorInfo := make(map[string]*SectorInfo)
	for _, stock := range stocks {
		sector := sectors[stock.Ticker]
		if _, exists := sectorInfo[sector]; !exists {
			sectorInfo[sector] = &SectorInfo{
				Value:  stock.Value,
				Stocks: []Stock{stock},
				Sector: sector,
			}
		} else {
			sectorInfo[sector].Stocks = append(sectorInfo[sector].Stocks, stock)
			sectorInfo[sector].Value += stock.Value
		}
	}

	portfolio := PortfolioValue{
		Sectors:          []SectorInfo{},
		TotalValue:       0,
		SectorPercentage: make(map[string]float32),
	}
	for _, sector := range sectorInfo {
		portfolio.TotalValue += sector.Value
		portfolio.Sectors = append(portfolio.Sectors, *sector)
	}
	for _, sector := range portfolio.Sectors {
		portfolio.SectorPercentage[sector.Sector] = (sector.Value / portfolio.TotalValue) * 100
	}
	return &portfolio
}
