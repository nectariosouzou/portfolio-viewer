// handlers/interface.go
package handlers

type GptHandler interface {
	FindSectors(tickers []string) (map[string]string, error)
}

type RedisHandler interface {
	FindSectors(tickers map[string]bool) (map[string]string, map[string]bool, error)
	SetTicker(tickers map[string]string) error
}
