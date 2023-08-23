package api

import (
	"reflect"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

var stocks = []Stock{
	{
		Name:   "Apple",
		Ticker: "AAPL",
		Price:  100,
		Value:  300,
		Shares: 3,
	},
	{
		Name:   "Citi",
		Ticker: "C",
		Price:  100,
		Value:  300,
		Shares: 3,
	},
	{
		Name:   "Alphabet",
		Ticker: "GOOG",
		Price:  100,
		Value:  300,
		Shares: 3,
	},
}

type testCase struct {
	Input               []Stock
	ExpectedGptInput    []string
	ExpectedRedisOutput RedisOutput
	ExpectedGptOutput   map[string]string
	Output              map[string]string
}

type RedisOutput struct {
	SectorTable map[string]string
	BoolTable   map[string]bool
}

func TestSortStocks(t *testing.T) {
	var testCases = []testCase{
		{
			Input: stocks,
			ExpectedRedisOutput: RedisOutput{
				SectorTable: map[string]string{
					"APPL": "IT",
				},
				BoolTable: map[string]bool{
					"APPL": true,
					"Citi": false,
					"GOOG": false,
				},
			},
			ExpectedGptInput: []string{
				"Citi", "GOOG",
			},
			ExpectedGptOutput: map[string]string{
				"Citi": "Financial",
				"GOOG": "IT",
			},
			Output: map[string]string{
				"APPL": "IT",
				"Citi": "Financial",
				"GOOG": "IT",
			},
		},
		{
			Input: stocks,
			ExpectedRedisOutput: RedisOutput{
				SectorTable: map[string]string{},
				BoolTable: map[string]bool{
					"APPL": false,
					"Citi": false,
					"GOOG": false,
				},
			},
			ExpectedGptInput: []string{
				"APPL", "Citi", "GOOG",
			},
			ExpectedGptOutput: map[string]string{
				"APPL": "IT",
				"Citi": "Financial",
				"GOOG": "IT",
			},
			Output: map[string]string{
				"APPL": "IT",
				"Citi": "Financial",
				"GOOG": "IT",
			},
		},

		{
			Input: []Stock{},
			ExpectedRedisOutput: RedisOutput{
				SectorTable: map[string]string{},
				BoolTable:   map[string]bool{},
			},
			ExpectedGptInput:  []string{},
			ExpectedGptOutput: map[string]string{},
			Output:            map[string]string{},
		},
	}

	for i := 0; i < len(testCases); i++ {
		ctrl := gomock.NewController(t)

		mockRedisHandler := NewMockRedisHandler(ctrl)
		mockGptHandler := NewMockGptHandler(ctrl)
		var api = Api{
			RedisHandler: mockRedisHandler,
			GptHandler:   mockGptHandler,
		}
		mockRedisHandler.EXPECT().FindSectors(gomock.Any()).Return(testCases[i].ExpectedRedisOutput.SectorTable, testCases[i].ExpectedRedisOutput.BoolTable, nil)
		if len(testCases[i].ExpectedGptInput) != 0 {
			mockGptHandler.EXPECT().FindSectors(testCases[i].ExpectedGptInput).Return(testCases[i].ExpectedGptOutput, nil)
			mockRedisHandler.EXPECT().SetTicker(testCases[i].ExpectedGptOutput).Return(nil)
		}
		result, _ := api.sortStocks(stocks)
		for j := 0; j < len(testCases[i].Output); j++ {
			eq := reflect.DeepEqual(result, testCases[i].Output)
			if !eq {
				t.Fail()
			}
		}

	}

}
