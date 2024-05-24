package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	database "github.com/vistastaking/staking-indexer/db"
	"github.com/vistastaking/staking-indexer/handlers"
	indexerPkg "github.com/vistastaking/staking-indexer/indexer"
)

func main() {
	_ = godotenv.Load()

	db, err := database.Setup(os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("https://lb.nodies.app/v1/eda527f40f4c48698a739e2dfae256b5")
	if err != nil {
		log.Fatal(err)
	}

	indexer := indexerPkg.NewIndexer(client, db)

	stop := indexer.Start(
		[]indexerPkg.ProcessLogsInput{
			{
				StartBlock:      19_796_144,
				ContractAddress: handlers.RocketMinipoolManagerAddress,
				EventSigHash:    crypto.Keccak256Hash([]byte("MinipoolCreated(address,address,uint256)")),
				Handler:         handlers.RocketPoolTVL,
			},
			{
				StartBlock:      19_796_144,
				ContractAddress: handlers.RocketMinipoolManagerAddress,
				EventSigHash:    crypto.Keccak256Hash([]byte("MinipoolCreated(address,address,uint256)")),
				Handler:         handlers.RocketPoolTVL,
			},
		},
	)
	<-stop
}
