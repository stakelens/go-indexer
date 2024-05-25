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

	database "github.com/vistastaking/staking-indexer/db"
)

type Handler func(HandlerParams)

type HandlerParams struct {
	Client   *ethclient.Client
	Database *database.Queries
	Log      types.Log
}

type ProcessLogsInRangeInput struct {
	ContractAddress common.Address
	EventSigHash    common.Hash
	StartBlock      uint64
	EndBlock        uint64
	Handler         Handler
}

type IntervalOptions struct {
	Fn       func()
	Interval time.Duration
}

type indexer struct {
	client   *ethclient.Client
	database *database.Queries
}

func NewIndexer(client *ethclient.Client, database *database.Queries) *indexer {
	return &indexer{
		client:   client,
		database: database,
	}
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
				fmt.Println("Stopping interval")
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
}

func (i *indexer) Start(inputs []ProcessLogsInput) chan bool {
	stopSignals := make([]chan bool, len(inputs))

	for index, input := range inputs {
		stopSignals[index] = i.ProcessLogsInRealTime(input)
	}

	stop := make(chan bool)

	go func() {
		<-stop
		for _, stopSignal := range stopSignals {
			stopSignal <- true
		}
	}()

	return stop
}

func (i *indexer) ProcessLogsInRealTime(input ProcessLogsInput) chan bool {
	var indexerRunning atomic.Bool

	startBlock := atomic.Uint64{}
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

			mostRecentBlockNumber, err := i.client.BlockNumber(context.Background())
			if err != nil {
				fmt.Println("Error getting most recent block number.")
				log.Fatal(err)
			}

			if mostRecentBlockNumber == startBlock.Load() {
				fmt.Println("Indexer is up to date. Skipping this interval.")
				indexerRunning.Store(false)
				return
			}

			i.processLogsInRange(ProcessLogsInRangeInput{
				StartBlock:      startBlock.Load(),
				EndBlock:        mostRecentBlockNumber,
				ContractAddress: input.ContractAddress,
				EventSigHash:    input.EventSigHash,
				Handler:         input.Handler,
			})

			fmt.Println("Finished processing logs.")
			indexerRunning.Store(false)
			startBlock.Store(mostRecentBlockNumber)
		},
		Interval: 1 * time.Second,
	})

	return stop
}

func (i *indexer) processLogsInRange(input ProcessLogsInRangeInput) {
	var currentBlock uint64 = input.StartBlock
	var stepSize uint64 = 10_000

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

		logs, err := filterLogsWithRetry(
			filterLogsWithRetryInput{
				maxRetries: 5,
				sleep:      1 * time.Second,
				context:    context.Background(),
				query:      query,
				client:     i.client,
			},
		)

		if err != nil {
			fmt.Println("Error filtering logs.")
			log.Fatal(err)
		}

		fmt.Printf("Amount of logs: %d\n", len(logs))

		for _, vLog := range logs {
			input.Handler(HandlerParams{
				Client:   i.client,
				Database: i.database,
				Log:      vLog,
			})
		}

		currentBlock = rangeEndBlock
	}
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
		fmt.Println("Error marshalling query")
		log.Fatal(err)
	}

	hash := sha256.New()
	hash.Write(queryBytes)
	return hex.EncodeToString(hash.Sum(nil))
}
