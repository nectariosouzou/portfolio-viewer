package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type GptClient struct {
	Client *openai.Client
}

type GptSector struct {
	Sector string `json:"Sector"`
	Ticker string `json:"Ticker"`
}

func InitGpthandler(key string) *GptClient {
	return &GptClient{Client: openai.NewClient(key)}
}

func (g *GptClient) FindSectors(tickers []string) (map[string]string, error) {
	str := message(tickers)
	resp, err := g.Client.CreateChatCompletion(
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
		return nil, err
	}
	classification := resp.Choices[0].Message.Content
	var data []GptSector
	err = json.Unmarshal([]byte(classification), &data)
	if err != nil {
		return nil, err
	}
	sectors := make(map[string]string)
	for _, item := range data {
		sectors[item.Ticker] = item.Sector
	}
	return sectors, nil
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
