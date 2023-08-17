package handlers

import (
	"context"
	"fmt"

	"github.com/redis/rueidis"
)

type RedisHandler struct {
	redisClient rueidis.Client
}

type Sector struct {
	Tickers []string
}

func InitRedisHandler(client rueidis.Client) RedisHandler {
	return RedisHandler{
		redisClient: client,
	}
}

func (r *RedisHandler) FindSectors(tickers map[string]bool) (map[string]string, error) {
	ctx := context.Background()
	cmds := make(rueidis.Commands, 0, len(tickers))
	for key := range tickers {
		cmds = append(cmds, r.redisClient.B().Hgetall().Key(key).Build())
	}
	sectors := map[string]string{}
	for _, resp := range r.redisClient.DoMulti(ctx, cmds...) {
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

func (r *RedisHandler) SetTicker(tickers map[string]string) error {
	ctx := context.Background()
	cmds := make(rueidis.Commands, 0, len(tickers))
	for key, value := range tickers {
		cmds = append(cmds, r.redisClient.B().Hset().Key(key).FieldValue().FieldValue("Ticker", key).FieldValue("Sector", value).Build())
	}
	for _, resp := range r.redisClient.DoMulti(ctx, cmds...) {
		if err := resp.Error(); err != nil {
			fmt.Printf("redis error: %s", err)
			return err
		}
	}
	return nil
}
