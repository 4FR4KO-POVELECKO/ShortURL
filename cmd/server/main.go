package main

import (
	"ShortURL/internal/app/store"
	"ShortURL/pkg/api"
	"ShortURL/pkg/grpcserver"
	"log"
	"net"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Get env
	redis_host, exists := os.LookupEnv("REDIS_HOST")
	if !exists {
		redis_host = "127.0.0.1"
	}

	redisPort, exists := os.LookupEnv("REDIS_PORT")
	if !exists {
		redisPort = "6379"
	}

	grpcPort, exists := os.LookupEnv("GRPC_PORT")
	if !exists {
		grpcPort = "5000"
	}

	// Connect Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redis_host + ":" + redisPort,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStoreRedis(client)

	// GRPC server
	s := grpc.NewServer()
	srv := &grpcserver.GRPCServer{Store: store}

	api.RegisterShortlinkServer(s, srv)

	l, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(s.Serve(l))
}
