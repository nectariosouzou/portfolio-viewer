package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type GptSector struct {
	Sector string `json:"Sector"`
	Ticker string `json:"Ticker"`
}

func message(tickers []string) string {
	arrStr := strings.Join(tickers, ",")
	str := fmt.Sprintf(`Classify these stock into the 11 GICS sectors and two custom sectors, which are 
	[
	"ENERGY",
	"MATERIALS",
	"INDUSTRIALS",
	"CONSUMER DISCRETIONARY",
	"CONSUMER STAPLES",
	"HEALTH CARE",
	"FINANCIALS",
	"INFORMATION TECHNOLOGY",
	"COMMUNICATION SERVICES",
	"UTILITIES",
	"REAL ESTATE",
	"BONDS AND EQ",
	]: %s. Only output as [{"Ticker": x, "Sector": y}, ...] No special characters, only letters and spaces`,
		arrStr)
	return str
}

func findSectors(tickers []string) (map[string]string, error) {
	str := message(tickers)
	client := openai.NewClient(Env["API_KEY"])
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: 0.5,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: str,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}
	classification := resp.Choices[0].Message.Content
	var data []GptSector
	err = json.Unmarshal([]byte(classification), &data)
	if err != nil {
		fmt.Println(fmt.Errorf("error with Unmarshall: %s", err))
		return nil, err
	}
	sectors := make(map[string]string)
	for _, item := range data {
		sectors[item.Ticker] = item.Sector
	}
	return sectors, nil
}
