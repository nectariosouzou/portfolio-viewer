package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/redis/rueidis"
)

var Env map[string]string

type DbClients struct {
	RedisClient rueidis.Client
}

func InitDbClients() *DbClients {
	Env = loadEnv()
	return &DbClients{
		RedisClient: NewRedisClient([]string{Env["REDIS_ADDR"]}),
	}
}

func NewRedisClient(addr []string) rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: addr})
	if err != nil {
		log.Fatalf("error starting redis client: %s", err)
	}
	return client
}

func loadEnv() map[string]string {
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return envs
}
