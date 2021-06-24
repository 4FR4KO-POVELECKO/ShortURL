package main

import (
	"ShortURL/internal/app/apiserver"
	"ShortURL/pkg/api"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewShortlinkClient(conn)

	apiserver.Start("localhost:8000", c)
}
