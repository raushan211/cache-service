package database

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const (
	prefix = "raushan211"
)

type redisDatabase struct {
	client *redis.Client
}

// CreateRedisDatabase creates the redis database
func createRedisDatabase() (Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result() // makes sure database is connected
	if err != nil {
		fmt.Println(err)
		return nil, &CreateDatabaseError{}
	}
	return &redisDatabase{client: client}, nil
}
func (r *redisDatabase) PrepareKey(key string) string {
	return fmt.Sprint(prefix, ":", key)
}

//Set ttl to 300 seconds
func (r *redisDatabase) Set(key string, value string) ([]byte, error) {
	_, err := r.client.Set(r.PrepareKey(key), value, time.Second*300).Result()
	if err != nil {
		return generateError("set", err)
	}
	return []byte(key), nil
}

func (r *redisDatabase) Get(key string) ([]byte, error) {
	value, err := r.client.Get(r.PrepareKey(key)).Bytes()
	if err != nil {
		return generateError("get", err)
	}
	return value, nil

}

func generateError(operation string, err error) ([]byte, error) {
	if err == redis.Nil {
		return nil, &OperationError{operation}
	}
	return nil, &DownError{}
}
