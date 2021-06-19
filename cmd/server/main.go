package main

import (
	"ShortURL/pkg/api"
	"ShortURL/pkg/shorturl"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &shorturl.GRPCServer{}
	api.RegisterShortlinkServer(s, srv)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
