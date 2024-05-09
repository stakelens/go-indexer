package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	database "github.com/vistastaking/staking-indexer/db"
	"github.com/vistastaking/staking-indexer/handlers"
	"github.com/vistastaking/staking-indexer/indexer"
)

func main() {
	_ = godotenv.Load()

	db, err := database.Setup(os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	stop := indexer.ProcessLogsInRealTime(
		indexer.ProcessLogsInput{
			StartBlock:      19_796_144,
			ContractAddress: handlers.RocketMinipoolManagerAddress,
			EventSigHash:    crypto.Keccak256Hash([]byte("MinipoolCreated(address,address,uint256)")),
			Handler:         handlers.RocketPoolTVL,
			Database:        db,
		},
	)
	<-stop
}
