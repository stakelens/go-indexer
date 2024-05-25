package main

import (
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/vistastaking/staking-indexer/db"
	"github.com/vistastaking/staking-indexer/handlers"
	indexerPkg "github.com/vistastaking/staking-indexer/indexer"
	"github.com/vistastaking/staking-indexer/rpc"
)

func main() {
	_ = godotenv.Load()
	rpc, err := rpc.New("localhost:8080" /*"postgres://localhost:5432" */, "https://lb.nodies.app/v1/eda527f40f4c48698a739e2dfae256b5")

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		rpc.Server.ListenAndServe()
	}()

	ctx := context.Background()
	pgx, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	dbQueries := db.New(pgx)

	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	indexer := indexerPkg.NewIndexer(client, dbQueries)

	stop := indexer.Start(
		[]indexerPkg.ProcessLogsInput{
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
