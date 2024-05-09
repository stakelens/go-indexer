package main

import (
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vistastaking/staking-indexer/database"
	"github.com/vistastaking/staking-indexer/handlers"
	"github.com/vistastaking/staking-indexer/indexer"
)

func main() {
	db, err := database.Setup("postgres://postgres:1234@localhost:5432/ti1")

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
