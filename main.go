package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	abis "github.com/vistastaking/custom-staking-indexer/abis"
)

var rocketVaultAddress = common.HexToAddress("0x3bdc69c4e5e13e52a65f5583c23efb9636b469d6")
var rocketNodeStakingAddress = common.HexToAddress("0x0d8d8f8541b12a0e1194b7cc4b6d954b90ab82ec")
var rocketMinipoolManagerAddress = common.HexToAddress("0x6d010c43d4e96d74c422f2e27370af48711b49bf")

func rocketPoolTVL(blockNumber *big.Int) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/9282a5f3ed9c41efa8c5176a8c644852")
	if err != nil {
		log.Fatal(err)
	}

	rocketVaultContract, err := abis.NewRocketVault(rocketVaultAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	rocketNodeStakingContract, err := abis.NewRocketNodeStaking(rocketNodeStakingAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	rocketMinipoolManagerContract, err := abis.NewRocketMinipoolManager(rocketMinipoolManagerAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	totalETH := big.NewInt(0)
	totalRPL := big.NewInt(0)

	limit := big.NewInt(1000)
	offset := big.NewInt(0)

	initialisedMinipools := big.NewInt(0)
	prelaunchMinipools := big.NewInt(0)
	stakingMinipools := big.NewInt(0)
	withdrawableMinipools := big.NewInt(0)

	for {
		activeMinipools, err := rocketMinipoolManagerContract.GetMinipoolCountPerStatus(&bind.CallOpts{BlockNumber: blockNumber}, offset, limit)
		if err != nil {
			log.Fatal(err)
		}

		initialisedMinipools.Add(initialisedMinipools, activeMinipools.InitialisedCount)
		prelaunchMinipools.Add(prelaunchMinipools, activeMinipools.PrelaunchCount)
		stakingMinipools.Add(stakingMinipools, activeMinipools.StakingCount)
		withdrawableMinipools.Add(withdrawableMinipools, activeMinipools.WithdrawableCount)

		// If the sum of all minipools is less than the limit, we have reached the end
		total := big.NewInt(0)
		total = total.Add(total, activeMinipools.InitialisedCount)
		total = total.Add(total, activeMinipools.PrelaunchCount)
		total = total.Add(total, activeMinipools.StakingCount)
		total = total.Add(total, activeMinipools.WithdrawableCount)
		total = total.Add(total, activeMinipools.DissolvedCount)

		if total.Cmp(limit) == -1 {
			break
		}

		offset.Add(offset, limit)
	}

	ethLockedInMinipools := big.NewInt(0)
	ethLockedInMinipools = ethLockedInMinipools.Add(ethLockedInMinipools, initialisedMinipools.Mul(initialisedMinipools, big.NewInt(16)))
	ethLockedInMinipools = ethLockedInMinipools.Add(ethLockedInMinipools, prelaunchMinipools.Mul(prelaunchMinipools, big.NewInt(32)))
	ethLockedInMinipools = ethLockedInMinipools.Add(ethLockedInMinipools, stakingMinipools.Mul(stakingMinipools, big.NewInt(32)))
	ethLockedInMinipools = ethLockedInMinipools.Add(ethLockedInMinipools, withdrawableMinipools.Mul(withdrawableMinipools, big.NewInt(32)))
	ethLockedInMinipools = ethLockedInMinipools.Mul(ethLockedInMinipools, big.NewInt(1e18))

	totalETH.Add(totalETH, ethLockedInMinipools)

	rocketDepositPoolEth, err := rocketVaultContract.BalanceOf(&bind.CallOpts{BlockNumber: blockNumber}, "rocketDepositPool")
	if err != nil {
		log.Fatal(err)
	}

	totalETH.Add(totalETH, rocketDepositPoolEth)

	totalRPLStaked, err := rocketNodeStakingContract.GetTotalRPLStake(&bind.CallOpts{BlockNumber: blockNumber})
	if err != nil {
		log.Fatal(err)
	}

	totalRPL.Add(totalRPL, totalRPLStaked)

	rocketDAONodeTrustedActionsRPLBalance, err := rocketVaultContract.BalanceOfToken(&bind.CallOpts{BlockNumber: blockNumber}, "rocketDAONodeTrustedActions", rocketVaultAddress)
	if err != nil {
		log.Fatal(err)
	}

	totalRPL.Add(totalRPL, rocketDAONodeTrustedActionsRPLBalance)

	rocketAuctionManagerRPLbalance, err := rocketVaultContract.BalanceOfToken(&bind.CallOpts{BlockNumber: blockNumber}, "rocketAuctionManager", rocketVaultAddress)
	if err != nil {
		log.Fatal(err)
	}

	totalRPL.Add(totalRPL, rocketAuctionManagerRPLbalance)

	fmt.Printf("Total ETH locked in minipools: %s\n", totalETH.String())
	fmt.Printf("Total RPL staked: %s\n", totalRPL.String())
}

func main() {
	rocketPoolTVL(big.NewInt(19520243))
}
