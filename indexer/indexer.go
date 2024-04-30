package indexer

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Handler func(types.Log)

type ProcessLogsInRangeInput struct {
	Client          *ethclient.Client
	ContractAddress common.Address
	EventSigHash    common.Hash
	StartBlock      uint64
	EndBlock        uint64
	Handler         Handler
}

func ProcessLogsInRange(input ProcessLogsInRangeInput) {
	var currentBlock uint64 = input.StartBlock
	var stepSize uint64 = 10_000

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
				client:     input.Client,
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		for _, vLog := range logs {
			input.Handler(vLog)
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
