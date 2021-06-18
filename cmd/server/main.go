package main

import (
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &shortlink.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
