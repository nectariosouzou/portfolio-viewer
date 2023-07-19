package main

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
)

type Stock struct {
	Name   string
	Ticker string
	Value  float32
	Price  float32
	Shares float32
}

func createStockList(data [][]string) []Stock {
	var stockList []Stock
	for i, line := range data {
		if i > 0 { // omit header line
			var rec Stock
			for j, field := range line {
				switch j {
				case 2:
					rec.Ticker = field
				case 3:
					runes := []rune(field)
					if len(runes) < 50 {
						rec.Name = field
					} else {
						rec.Name = string(runes[:50])
					}
				case 4:
					noDollar := strings.Replace(field, "$", "", -1)
					val, err := strconv.ParseFloat(noDollar, 32)
					if err != nil {
						fmt.Println(err)
					} else {
						rec.Shares = float32(val)
					}
				case 5:
					noDollar := strings.Replace(field, "$", "", -1)
					val, err := strconv.ParseFloat(noDollar, 32)
					if err != nil {
						fmt.Println(err)
					} else {
						rec.Price = float32(val)
					}
				case 7:
					noDollar := strings.Replace(field, "$", "", -1)
					val, err := strconv.ParseFloat(noDollar, 32)
					if err != nil {
						fmt.Println(err)
					} else {
						rec.Value = float32(val)
					}
				}
			}
			stockList = append(stockList, rec)
		}
	}
	return stockList
}

func processStocks(input multipart.File) ([]Stock, error) {
	// read csv values using csv.Reader
	csvReader := csv.NewReader(input)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	// convert records to array of structs
	stockList := createStockList(data)
	return stockList, nil
}
