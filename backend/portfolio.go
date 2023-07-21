package main

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

func calculatePortfolio(stocks []Stock, sectors map[string]string) *PortfolioValue {
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

func sortStocks(stocks []Stock) (map[string]string, error) {
	tickers := make(map[string]bool, len(stocks))
	var tickersForGPT []string
	var sectors map[string]string
	for _, stock := range stocks {
		tickers[stock.Ticker] = false
	}
	sectors, err := db.findSectors(tickers)
	if err != nil {
		return nil, err
	}
	for key, ticker := range tickers {
		if !ticker {
			tickersForGPT = append(tickersForGPT, key)
		}
	}
	if len(tickersForGPT) > 0 {
		sectorsGPT, err := findSectors(tickersForGPT)
		if err != nil {
			return nil, err
		}
		err = db.SetTicker(sectors)
		if err != nil {
			return nil, err
		}
		for key, sector := range sectorsGPT {
			sectors[key] = sector
		}
	}
	return sectors, nil
}
