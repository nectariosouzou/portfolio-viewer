package main

import (
	"log"
	"os"

	"github.com/redis/rueidis"
)

var Env map[string]string

type DbClients struct {
	RedisClient rueidis.Client
}

func InitDbClients() *DbClients {
	return &DbClients{
		RedisClient: NewRedisClient([]string{os.Getenv("REDIS_ADDR")}),
	}
}

func NewRedisClient(addr []string) rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: addr})
	if err != nil {
		log.Fatalf("error starting redis client: %s", err)
	}
	return client
}

func InitEnv() {
	Env = make(map[string]string)
	Env["REDIS_ADDR"] = os.Getenv("REDIS_ADDR")
	Env["API_KEY"] = os.Getenv("API_KEY")
}
