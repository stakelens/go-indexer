package indexer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vistastaking/custom-staking-indexer/cache"
	"github.com/vistastaking/custom-staking-indexer/database"
)

type Handler func(HandlerParams)

type HandlerParams struct {
	Client   *ethclient.Client
	Database *database.Queries
	Log      types.Log
	Cache    *cache.Cache
}

type ProcessLogsInRangeInput struct {
	Client          *ethclient.Client
	ContractAddress common.Address
	EventSigHash    common.Hash
	StartBlock      uint64
	EndBlock        uint64
	Handler         Handler
	Database        *database.Queries
}

func ProcessLogsInRange(input ProcessLogsInRangeInput) {
	var currentBlock uint64 = input.StartBlock
	var stepSize uint64 = 10_000
	var cache = cache.NewCache(input.Database)

	for currentBlock < input.EndBlock {
		rangeEndBlock := currentBlock + stepSize

		if rangeEndBlock > input.EndBlock {
			rangeEndBlock = input.EndBlock
		}

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(currentBlock)),
			ToBlock:   big.NewInt(int64(rangeEndBlock)),
			Addresses: []common.Address{
				input.ContractAddress,
			},
			Topics: [][]common.Hash{
				{input.EventSigHash},
			},
		}

		logs := getLogsWithCache(
			GetLogsWithCacheInput{
				context: context.Background(),
				query:   query,
				cache:   cache,
				client:  input.Client,
			})

		for _, vLog := range logs {
			input.Handler(HandlerParams{
				Client:   input.Client,
				Database: input.Database,
				Log:      vLog,
				Cache:    cache,
			})
		}

		currentBlock = rangeEndBlock
	}
}

type GetLogsWithCacheInput struct {
	context context.Context
	query   ethereum.FilterQuery
	cache   *cache.Cache
	client  *ethclient.Client
}

func getLogsWithCache(input GetLogsWithCacheInput) []types.Log {
	queryHash := GetQueryHash(input.query)
	cachedLogs, err := input.cache.Get(queryHash)

	// If the logs are cached, return them
	if err == nil {
		var logs []types.Log
		err = json.Unmarshal([]byte(cachedLogs), &logs)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Fetched logs from cache. Query: %s\n", queryHash)
		return logs
	}

	// Fetch the logs from the client if they are not cached
	result, err := filterLogsWithRetry(
		filterLogsWithRetryInput{
			maxRetries: 5,
			sleep:      1 * time.Second,
			context:    context.Background(),
			query:      input.query,
			client:     input.client,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	// Cache the logs
	input.cache.Set(queryHash, result)
	return result
}

type filterLogsWithRetryInput struct {
	maxRetries int
	sleep      time.Duration
	context    context.Context
	query      ethereum.FilterQuery
	client     *ethclient.Client
}

func filterLogsWithRetry(config filterLogsWithRetryInput) ([]types.Log, error) {
	var err error
	var result []types.Log

	for i := 0; i < config.maxRetries; i++ {
		result, err = config.client.FilterLogs(config.context, config.query)

		if err == nil {
			return result, nil
		}

		time.Sleep(config.sleep)
		log.Printf("Retrying FilterLogs. Attempt %d\n", i+1)
	}

	return nil, err
}

func GetQueryHash(query ethereum.FilterQuery) string {
	queryBytes, err := json.Marshal(query)
	if err != nil {
		log.Fatal(err)
	}

	hash := sha256.New()
	hash.Write(queryBytes)
	return hex.EncodeToString(hash.Sum(nil))
}
