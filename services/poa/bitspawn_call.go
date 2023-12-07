/*

 */

package poa

import (
	"context"
	"errors"
	"github.com/bitspawngg/bitspawn-api/enum"
	"math"
	"math/big"

	"github.com/bitspawngg/bitspawn-api/contracts/pool"
	"github.com/ethereum/go-ethereum/common"
)

func (bpc *BitspawnPoaClient) GetSPWNBalance(userAddress string, feeType string) (*big.Float, error) {
	client, err := bpc.getClient()
	if err != nil {
		return nil, err
	}
	var tokenContract common.Address
	switch feeType {
	case enum.Credit.ToString():
		tokenContract = bpc.credContract
	case enum.Spwn.ToString():
		tokenContract = bpc.spwnContract
	case enum.Usdc.ToString():
		tokenContract = bpc.usdcContract
	default:
		return nil, errors.New("unknown fee type")
	}
	bitspawn, err := pool.NewBitspawn(tokenContract, client)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	balance, err := bitspawn.BalanceOf(nil, common.HexToAddress(userAddress))
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	return convertWeiToEth(balance), nil
}

func (bpc *BitspawnPoaClient) GetEthBalance(userAddress string) (*big.Float, error) {
	client, err := bpc.getClient()
	if err != nil {
		return nil, err
	}

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(userAddress), nil)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	return convertWeiToEth(balance), nil
}

func (bpc *BitspawnPoaClient) GetParticipantCount(tournamentAddress string) (int64, error) {
	client, err := bpc.getClient()
	if err != nil {
		return 0, err
	}

	tourney, err := pool.NewTournamentPool(common.HexToAddress(tournamentAddress), client)
	if err != nil {
		bpc.log.Error(err)
		return 0, err
	}

	participantsCount, err := tourney.NumPlayers(nil)
	if err != nil {
		bpc.log.Error(err)
		return 0, err
	}

	return participantsCount.Int64(), nil
}

func (bpc *BitspawnPoaClient) GetPrizePool(tournamentAddress, feeType string) (int64, error) {
	client, err := bpc.getClient()
	if err != nil {
		return 0, err
	}

	var tokenContract common.Address
	switch feeType {
	case enum.Credit.ToString():
		tokenContract = bpc.credContract
	case enum.Spwn.ToString():
		tokenContract = bpc.spwnContract
	case enum.Spwn.ToString():
		tokenContract = bpc.usdcContract
	default:
		return 0, errors.New("unknown fee type")
	}

	bitspawn, err := pool.NewBitspawn(tokenContract, client)
	if err != nil {
		bpc.log.Error(err)
		return 0, err
	}

	prizePool, err := bitspawn.BalanceOf(nil, common.HexToAddress(tournamentAddress))
	if err != nil {
		bpc.log.Error(err)
		return 0, err
	}

	prizePoolInt, _ := convertWeiToEth(prizePool).Int64()
	return prizePoolInt, nil
}

func (bpc *BitspawnPoaClient) GetTournamentStatus(tournamentAddress string) (string, error) {
	client, err := bpc.getClient()
	if err != nil {
		return "", err
	}

	tournament, err := pool.NewTournamentPool(common.HexToAddress(tournamentAddress), client)
	if err != nil {
		bpc.log.Error(err)
		return "", err
	}

	status, err := tournament.IsCompleted(nil)
	if err != nil {
		bpc.log.Error(err)
		return "", err
	}

	switch status {
	case true:
		return "Completed", nil
	case false:
		return "Ongoing", nil
	default:
		return "", errors.New("unknown status")
	}
}

func (bpc *BitspawnPoaClient) IsRegistered(userAddress, tournamentAddress string) (bool, error) {
	client, err := bpc.getClient()
	if err != nil {
		return false, err
	}

	tourney, err := pool.NewTournamentPool(common.HexToAddress(tournamentAddress), client)
	if err != nil {
		bpc.log.Error(err)
		return false, err
	}

	isRegistered, err := tourney.Players(nil, common.HexToAddress(userAddress))
	if err != nil {
		bpc.log.Error(err)
		return false, err
	}

	return isRegistered, nil
}

func convertWeiToEth(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetString(wei.String())
	ethValue := new(big.Float).Quo(f, big.NewFloat(math.Pow10(18)))
	return ethValue
}
