package main

import (
	"ShortURL/internal/app/apiserver"
	"ShortURL/pkg/api"
	"log"
	"os"

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
	host, exists := os.LookupEnv("HOST")
	if !exists {
		host = "localhost"
	}

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8000"
	}

	grpcPort, exists := os.LookupEnv("GRPC_PORT")
	if !exists {
		grpcPort = "5000"
	}

	// Connet to grpc server
	conn, err := grpc.Dial(":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewShortlinkClient(conn)

	// Start api server
	apiserver.Start(host+":"+port, c)
}
