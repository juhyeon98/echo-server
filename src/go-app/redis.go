package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASS")

	if redisHost == "" || redisPort == "" {
		log.Fatalln("REDIS_HOST or REDIS_PORT enviroment are not set")
	}
	address := fmt.Sprintf("%s:%s", redisHost, redisPort)
	rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: redisPass,
		DB:       0,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalln("Fail to connect Redis")
	}
	fmt.Println("Success to connect Redis")
}

func loggingRedis(address *net.UDPAddr) {
	err := rdb.Set(address.IP.String(), "true", 0).Err()
	if err != nil {
		log.Printf("Fail to set %s", address.IP.String())
	}
}

func closeRedis() {
	rdb.Close()
}
