package cache

import (
	"context"
	"encoding/json"
	"log"

	"github.com/vistastaking/custom-staking-indexer/database"
)

type Cache struct {
	Database *database.Queries
}

func NewCache(database *database.Queries) *Cache {
	return &Cache{database}
}

func (cache *Cache) Set(key string, value any) {
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
	result, err := cache.Database.GetCache(context.Background(), key)
	if err != nil {
		return "", err
	}

	return result.Data, nil
}
