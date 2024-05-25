package rpc

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	database "github.com/vistastaking/staking-indexer/db"
)

type Cache struct {
	Database *database.Queries
	mu       sync.RWMutex
}

func NewCache(database *database.Queries) *Cache {
	return &Cache{Database: database}
}

func (cache *Cache) Set(key string, value any) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	result, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}

	cache.Database.Cache(context.Background(), database.CacheParams{
		ID:   key,
		Data: string(result),
	})
}

func (cache *Cache) Get(key string) (string, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	result, err := cache.Database.GetCache(context.Background(), key)
	if err != nil {
		return "", err
	}

	return result.Data, nil
}
