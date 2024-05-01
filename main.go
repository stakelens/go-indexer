package main

import abigeb "github.com/vistastaking/custom-staking-indexer/lib/abigen"

// import (
// 	"context"
// 	"log"

// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/vistastaking/custom-staking-indexer/database"
// 	"github.com/vistastaking/custom-staking-indexer/handlers"
// 	"github.com/vistastaking/custom-staking-indexer/indexer"
// )

func main() {
	abigeb.Abigen(abigeb.ABIGenOptions{
		Abi:    "./storage/rocket.abi",
		Pkg:    "storage",
		Type:   "Rocket",
		Output: "./storage/rocket.go",
	})

	// db, err := database.Setup()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client, err := ethclient.Dial("https://lb.nodies.app/v1/eda527f40f4c48698a739e2dfae256b5")
	// // client, err := ethclient.Dial("https://mainnet.infura.io/v3/9282a5f3ed9c41efa8c5176a8c644852")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// mostRecentBlockNumber, err := client.BlockNumber(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// indexer.ProcessLogsInRange(
	// 	indexer.ProcessLogsInRangeInput{
	// 		Client:          client,
	// 		StartBlock:      19_522_200,
	// 		EndBlock:        mostRecentBlockNumber,
	// 		ContractAddress: handlers.RocketMinipoolManagerAddress,
	// 		EventSigHash:    crypto.Keccak256Hash([]byte("MinipoolCreated(address,address,uint256)")),
	// 		Handler:         handlers.RocketPoolTVL,
	// 		Database:        db,
	// 	},
	// )
}
