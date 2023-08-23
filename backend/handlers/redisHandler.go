package handlers

import (
	"context"

	"github.com/redis/rueidis"
)

type RedisClient struct {
	redisClient rueidis.Client
}

type Sector struct {
	Tickers []string
}

func InitRedisHandler(client rueidis.Client) *RedisClient {
	return &RedisClient{
		redisClient: client,
	}
}

func (r *RedisClient) FindSectors(tickers map[string]bool) (map[string]string, map[string]bool, error) {
	ctx := context.Background()
	cmds := make(rueidis.Commands, 0, len(tickers))
	for key := range tickers {
		cmds = append(cmds, r.redisClient.B().Hgetall().Key(key).Build())
	}
	sectors := map[string]string{}
	for _, resp := range r.redisClient.DoMulti(ctx, cmds...) {
		if err := resp.Error(); err != nil {
			return nil, nil, err
		}
		mp, err := resp.AsStrMap()
		if err != nil {
			return nil, nil, err
		}
		tickers[mp["Ticker"]] = true
		sectors[mp["Ticker"]] = mp["Sector"]
	}
	return sectors, tickers, nil
}

func (r *RedisClient) SetTicker(tickers map[string]string) error {
	ctx := context.Background()
	cmds := make(rueidis.Commands, 0, len(tickers))
	for key, value := range tickers {
		cmds = append(cmds, r.redisClient.B().Hset().Key(key).FieldValue().FieldValue("Ticker", key).FieldValue("Sector", value).Build())
	}
	for _, resp := range r.redisClient.DoMulti(ctx, cmds...) {
		if err := resp.Error(); err != nil {
			return err
		}
	}
	return nil
}
