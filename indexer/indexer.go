package indexer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vistastaking/staking-indexer/cache"
	"github.com/vistastaking/staking-indexer/database"
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

type IntervalOptions struct {
	Fn       func()
	Interval time.Duration
}

func SetInterval(options IntervalOptions) chan bool {
	ticker := time.NewTicker(options.Interval)
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				options.Fn()
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

	return stop
}

type ProcessLogsInput struct {
	ContractAddress common.Address
	EventSigHash    common.Hash
	StartBlock      uint64
	Handler         Handler
	Database        *database.Queries
}

func ProcessLogsInRealTime(input ProcessLogsInput) chan bool {
	client, err := ethclient.Dial("https://lb.nodies.app/v1/eda527f40f4c48698a739e2dfae256b5")
	// client, err := ethclient.Dial("https://mainnet.infura.io/v3/9282a5f3ed9c41efa8c5176a8c644852")

	if err != nil {
		log.Fatal(err)
	}

	var indexerRunning atomic.Bool
	var startBlock atomic.Uint64 = atomic.Uint64{}

	startBlock.Store(input.StartBlock)
	indexerRunning.Store(false)

	stop := SetInterval(IntervalOptions{
		Fn: func() {
			fmt.Println("Checking for new logs...")

			if indexerRunning.Load() {
				fmt.Println("Indexer is already running. Skipping this interval.")
				return
			}

			indexerRunning.Store(true)

			mostRecentBlockNumber, err := client.BlockNumber(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			if mostRecentBlockNumber == startBlock.Load() {
				fmt.Println("Indexer is up to date. Skipping this interval.")
				indexerRunning.Store(false)
				return
			}

			ProcessLogsInRange(ProcessLogsInRangeInput{
				Client:          client,
				StartBlock:      startBlock.Load(),
				EndBlock:        mostRecentBlockNumber,
				ContractAddress: input.ContractAddress,
				EventSigHash:    input.EventSigHash,
				Handler:         input.Handler,
				Database:        input.Database,
			})

			fmt.Println("Finished processing logs.")
			indexerRunning.Store(false)
			startBlock.Store(mostRecentBlockNumber)
		},
		Interval: 1 * time.Second,
	})

	return stop
}

func ProcessLogsInRange(input ProcessLogsInRangeInput) {
	var currentBlock uint64 = input.StartBlock
	var stepSize uint64 = 10_000
	var cache = cache.NewCache(input.Database)

	fmt.Printf("Processing logs from block %d to %d\n", input.StartBlock, input.EndBlock)

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

		fmt.Printf("Amount of logs: %d\n", len(logs))

		// var wg sync.WaitGroup
		for _, vLog := range logs {
			// wg.Add(1)

			// go func(vLog types.Log) {
			input.Handler(HandlerParams{
				Client:   input.Client,
				Database: input.Database,
				Log:      vLog,
				Cache:    cache,
			})
			// 	wg.Done()
			// }(vLog)
		}

		// wg.Wait()
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
