package main

import (
	"context"
	"fmt"

	"github.com/redis/rueidis"
)

type Sector struct {
	Tickers []string
}

func (r *DbClients) findSectors(tickers map[string]bool) (map[string]string, error) {
	ctx := context.Background()
	cmds := make(rueidis.Commands, 0, len(tickers))
	for key := range tickers {
		cmds = append(cmds, r.RedisClient.B().Hgetall().Key(key).Build())
	}
	sectors := map[string]string{}
	for _, resp := range r.RedisClient.DoMulti(ctx, cmds...) {
		if err := resp.Error(); err != nil {
			fmt.Printf("redis error: %s", err)
			return nil, err
		}
		mp, err := resp.AsStrMap()
		if err != nil {
			return nil, err
		}
		tickers[mp["Ticker"]] = true
		sectors[mp["Ticker"]] = mp["Sector"]
	}
	return sectors, nil
}

func (r *DbClients) SetTicker(tickers map[string]string) error {
	ctx := context.Background()
	cmds := make(rueidis.Commands, 0, len(tickers))
	for key, value := range tickers {
		cmds = append(cmds, r.RedisClient.B().Hset().Key(key).FieldValue().FieldValue("Ticker", key).FieldValue("Sector", value).Build())
	}
	for _, resp := range r.RedisClient.DoMulti(ctx, cmds...) {
		if err := resp.Error(); err != nil {
			fmt.Printf("redis error: %s", err)
			return err
		}
	}
	return nil
}
