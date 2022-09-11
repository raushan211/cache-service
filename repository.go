package main

import "context"

type Data interface {
	// Function to set the data in cache using key and value.
	Set(ctx context.Context, key string, value []byte) error

	// Function to get the byte array data based on a given key
	Get(ctx context.Context, key string) ([]byte, error)
}
