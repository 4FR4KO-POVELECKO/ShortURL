package grpcserver

import (
	"context"

	shorten "ShortURL/internal/app/utils"
	"ShortURL/pkg/api"
)

// GRPCServer ...
type GRPCServer struct{}

func (s *GRPCServer) Create(ctx context.Context, req *api.OriginUrl) (*api.ShortUrl, error) {
	short := shorten.Shorten()

	return &api.ShortUrl{Url: short}, nil
}

func (s *GRPCServer) Get(context.Context, *api.ShortUrl) (*api.OriginUrl, error) {
	return &api.OriginUrl{Url: "origin"}, nil
}
