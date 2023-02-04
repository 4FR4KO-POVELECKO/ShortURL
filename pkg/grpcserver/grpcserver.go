package grpcserver

import (
	"context"

	"ShortURL/internal/app/model"
	"ShortURL/internal/app/store"
	shorten "ShortURL/internal/app/utils"
	"ShortURL/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {
	Store store.Store
}

// Create функция принимает оригинальный URL и генерирует токен сокрашенного
func (s *GRPCServer) Create(ctx context.Context, req *api.OriginUrl) (*api.ShortUrl, error) {
	// Генерация токена
	short := shorten.Shorten()

	// Валидация
	url := &model.URL{
		OriginURL: req.Url,
		ShortURL:  short,
	}

	err := url.ValidateURL()
	if err != nil {
		return nil, err
	}

	// url.OriginURL = shorten.AddHTTP(url.OriginURL)

	// Сохраняем в бд
	err = s.Store.Set(url.ShortURL, url.OriginURL, 0)
	if err != nil {
		return nil, err
	}

	return &api.ShortUrl{Url: short}, nil
}

// Get функция принимает соркашенный URL и возвращает оригинальный
func (s *GRPCServer) Get(ctx context.Context, req *api.ShortUrl) (*api.OriginUrl, error) {
	// Валидация токена
	url := &model.URL{
		ShortURL: req.Url,
	}

	err := url.ValidateShortURL()
	if err != nil {
		return nil, err
	}

	// Записываем получение ссылки для статы
	_, err = s.Store.Incr("stat_" + url.ShortURL)
	if err != nil {
		return nil, err
	}

	// Проверяем в бд
	val, err := s.Store.Get(url.ShortURL)
	if err != nil {
		return nil, err
	}

	return &api.OriginUrl{Url: val}, nil
}
