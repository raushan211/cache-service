package service

import (
	"context"
	"errors"
	"log"
)

type Repository interface {
	Set(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
}

type RedisDB struct {
	Log log.Logger
}

func NewRedisDB(logger *log.Logger) *RedisDB {
	return &RedisDB{Log: *logger}
}

func (r *RedisDB) Set(ctx context.Context, key string, value []byte) error {
	return errors.New("not implemented yet. Raushan will implement me")
}

func (r *RedisDB) Get(ctx context.Context, key string) ([]byte, error) {
	return []byte("<yourname> will implement me"), nil
}
