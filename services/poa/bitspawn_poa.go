/*

 */

package poa

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/bitspawngg/bitspawn-api/enum"
	"math/big"
	"time"

	"github.com/bitspawngg/bitspawn-api/services/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	"github.com/bitspawngg/bitspawn-api/contracts/pool"
	"github.com/bitspawngg/bitspawn-api/services"
)

type BitspawnPoaClient struct {
	url          string
	log          *logrus.Entry
	faucetKey    string
	spwnContract common.Address
	spwnFee      common.Address
	credContract common.Address
	credFee      common.Address
	usdcContract common.Address
	usdcFee      common.Address
	walletUrl    string
	client       *ethclient.Client
}

func NewBitspawnPoaClient(log *logrus.Logger, configService *config.Service) *BitspawnPoaClient {
	return &BitspawnPoaClient{
		url:          configService.GetConfig().RpcNodeUrl,
		log:          log.WithField("services", "bitspawn_poa"),
		faucetKey:    configService.GetConfig().FaucetKey,
		spwnContract: common.HexToAddress(configService.GetConfig().SpwnContract),
		spwnFee:      common.HexToAddress(configService.GetConfig().SpwnFee),
		credContract: common.HexToAddress(configService.GetConfig().CredContract),
		credFee:      common.HexToAddress(configService.GetConfig().CredFee),
		usdcContract: common.HexToAddress(configService.GetConfig().UsdcContract),
		usdcFee:      common.HexToAddress(configService.GetConfig().UsdcFee),
		walletUrl:    configService.GetConfig().WalletUrl,
	}
}

func (bpc *BitspawnPoaClient) getClient() (*ethclient.Client, error) {
	if bpc.client == nil {
		client, err := ethclient.Dial(bpc.url)
		if err != nil {
			bpc.log.Error(err)
			return nil, err
		}
		bpc.client = client
	}
	return bpc.client, nil
}

func (bpc *BitspawnPoaClient) DeployTournament(
	privKeyHex, feeType string, feePercent, organizerPercent int64,
	entryFee *big.Int, initialParticipants []common.Address,
) (
	common.Address, *types.Transaction, *pool.TournamentPool, error) {

	client, err := bpc.getClient()
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	contractFee := big.NewInt(feePercent)        // 5%
	organizerFee := big.NewInt(organizerPercent) // 10%

	auth, fromAddress, err := bpc.getAuth(privKeyHex, client, services.NewTournamentGasLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	// address : organizer address
	// input := "1.0"
	var tokenContract, feeContract common.Address
	switch feeType {
	case enum.Credit.ToString():
		tokenContract = bpc.credContract
		feeContract = bpc.credFee
	case enum.Spwn.ToString():
		tokenContract = bpc.spwnContract
		feeContract = bpc.spwnFee
	case enum.Usdc.ToString():
		tokenContract = bpc.usdcContract
		feeContract = bpc.usdcFee
	default:
		return common.Address{}, nil, nil, errors.New("unknown fee type")
	}
	contractAddress, tx, instance, err := pool.DeployTournamentPool(auth, client, tokenContract, feeContract,
		contractFee, organizerFee, fromAddress, entryFee, initialParticipants)

	if err != nil {
		bpc.log.Error(err)
		return common.Address{}, nil, nil, err
	}

	bpc.log.Debugf("Transaction: %#v \n", tx)

	ok, err := WaitTxSuccess(client, tx, time.Second*20, time.Second)
	if !ok {
		return common.Address{}, nil, nil, errors.New("transaction failed")
	}
	if err != nil {
		return common.Address{}, nil, nil, errors.New("tx receipt timed out")
	}

	return contractAddress, tx, instance, nil
}

// func (bpc *BitspawnPoaClient) RegisterTeam(privKeyHex string, contractAddress common.Address, feeInWei *big.Int) (*types.Transaction, error) {
// 	client, err := bpc.getClient()
// 	if err != nil {
// 		return nil, err
// 	}

// 	tourney, err := tournament.NewTournament(contractAddress, client)
// 	if err != nil {
// 		bpc.log.Error(err)
// 		return nil, err
// 	}

// 	auth, _, err := bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	bitspawn, err := tournament.NewBitspawn(bpc.spwnContract, client)
// 	tx, err := bitspawn.Approve(auth, contractAddress, feeInWei)
// 	if err != nil {
// 		bpc.log.Error(err)
// 		return nil, err
// 	}

// 	ok, err := WaitTxSuccess(client, tx, time.Second*20, time.Second)
// 	if !ok {
// 		return tx, errors.New("approve transaction failed")
// 	}
// 	if err != nil {
// 		return tx, errors.New("approve tx receipt timed out")
// 	}

// 	auth, _, err = bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tx, err = tourney.RegisterTeam(auth, feeInWei)
// 	bpc.log.Debugf("Transaction: %#v \n", tx)

// 	if err != nil {
// 		bpc.log.Error(err)
// 		return nil, err
// 	}

// 	ok, err = WaitTxSuccess(client, tx, time.Second*20, time.Second)
// 	if !ok {
// 		return tx, errors.New("registerTeam transaction failed")
// 	}
// 	if err != nil {
// 		return tx, errors.New("registerTeam tx receipt timed out")
// 	}

// 	return tx, nil
// }

func (bpc *BitspawnPoaClient) RegisterTournament(privKeyHex, feeType string, contractAddress common.Address, feeInWei *big.Int) (*types.Transaction, error) {
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

	tourney, err := pool.NewTournamentPool(contractAddress, client)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	auth, _, err := bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
	if err != nil {
		return nil, err
	}

	bitspawn, _ := pool.NewBitspawn(tokenContract, client)
	tx, err := bitspawn.Approve(auth, contractAddress, feeInWei)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	ok, err := WaitTxSuccess(client, tx, time.Second*20, time.Second)
	if !ok {
		return tx, errors.New("approve transaction failed")
	}
	if err != nil {
		return tx, errors.New("approve tx receipt timed out")
	}

	auth, _, err = bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
	if err != nil {
		return nil, err
	}

	tx, err = tourney.Register(auth, feeInWei)
	bpc.log.Debugf("Transaction: %#v \n", tx)

	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	ok, err = WaitTxSuccess(client, tx, time.Second*20, time.Second)
	if !ok {
		return tx, errors.New("register transaction failed")
	}
	if err != nil {
		return tx, errors.New("register tx receipt timed out")
	}

	return tx, nil
}

// func (bpc *BitspawnPoaClient) UnregisterTeam(privKeyHex string, contractAddress common.Address) (*types.Transaction, error) {
// 	client, err := bpc.getClient()
// 	if err != nil {
// 		return nil, err
// 	}

// 	tourney, err := tournament.NewTournament(contractAddress, client)
// 	if err != nil {
// 		bpc.log.Error(err)
// 		return nil, err
// 	}

// 	auth, _, err := bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tx, err := tourney.UnregisterTeam(auth)
// 	bpc.log.Debugf("Transaction: %#v \n", tx)

// 	if err != nil {
// 		bpc.log.Error(err)
// 		return nil, err
// 	}

// 	ok, err := WaitTxSuccess(client, tx, time.Second*20, time.Second)
// 	if !ok {
// 		return tx, errors.New("unregisterTeam transaction failed")
// 	}
// 	if err != nil {
// 		return tx, errors.New("unregisterTeam tx receipt timed out")
// 	}

// 	return tx, nil
// }

func (bpc *BitspawnPoaClient) UnregisterTournament(privKeyHex string, contractAddress common.Address) (*types.Transaction, error) {
	client, err := bpc.getClient()
	if err != nil {
		return nil, err
	}

	tourney, err := pool.NewTournamentPool(contractAddress, client)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	auth, _, err := bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
	if err != nil {
		return nil, err
	}

	tx, err := tourney.Unregister(auth)
	bpc.log.Debugf("Transaction: %#v \n", tx)

	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	ok, err := WaitTxSuccess(client, tx, time.Second*20, time.Second)
	if !ok {
		return tx, errors.New("unregister transaction failed")
	}
	if err != nil {
		return tx, errors.New("unregister tx receipt timed out")
	}

	return tx, nil
}

func (bpc *BitspawnPoaClient) FundTournament(privKeyHex, feeType string, contractAddress common.Address, fundsInWei *big.Int) (*types.Transaction, error) {
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

	auth, _, err := bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
	if err != nil {
		return nil, err
	}

	tourney, err := pool.NewTournamentPool(contractAddress, client)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	bitspawn, err := pool.NewBitspawn(tokenContract, client)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	// approvalFunds := fundsInWei.Mul(fundsInWei, big.NewInt(3))
	tx, err := bitspawn.Approve(auth, contractAddress, fundsInWei)
	if err != nil {
		bpc.log.Error(err)
		return nil, err

	}
	_, _ = WaitTxSuccess(client, tx, time.Second*20, time.Second)

	auth, _, err = bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
	if err != nil {
		return nil, err
	}

	tx, err = tourney.Fund(auth, fundsInWei)
	bpc.log.Debugf("Transaction: %#v \n", tx)

	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	ok, err := WaitTxSuccess(client, tx, time.Second*20, time.Second)
	if !ok {
		return tx, errors.New("transaction failed")
	}
	if err != nil {
		return tx, errors.New("tx receipt timed out")
	}

	return tx, nil
}

func (bpc *BitspawnPoaClient) CompleteTournament(privKeyHex string, contractAddress common.Address,
	finalPlacements []common.Address, prizePerTenThousand []*big.Int) (*types.Transaction, error) {

	client, err := bpc.getClient()
	if err != nil {
		return nil, err
	}

	tourney, err := pool.NewTournamentPool(contractAddress, client)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	auth, _, err := bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
	if err != nil {
		return nil, err
	}

	tx, err := tourney.CompleteTournament(auth, finalPlacements, prizePerTenThousand)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	ok, err := WaitTxSuccess(client, tx, time.Second*20, time.Second)
	if !ok {
		return tx, errors.New("transaction failed")
	}
	if err != nil {
		return tx, errors.New("tx receipt timed out")
	}

	return tx, nil
}

func (bpc *BitspawnPoaClient) CancelTournament(privKeyHex string, contractAddress common.Address) (*types.Transaction, error) {

	client, err := bpc.getClient()
	if err != nil {
		return nil, err
	}

	tourney, err := pool.NewTournamentPool(contractAddress, client)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	auth, _, err := bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
	if err != nil {
		return nil, err
	}

	tx, err := tourney.CancelTournament(auth)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	ok, err := WaitTxSuccess(client, tx, time.Second*20, time.Second)
	if !ok {
		return tx, errors.New("transaction failed")
	}
	if err != nil {
		return tx, errors.New("tx receipt timed out")
	}

	return tx, nil
}

func WaitTxSuccess(client *ethclient.Client, tx *types.Transaction, maxWait time.Duration,
	interval time.Duration) (bool, error) {

	ticker := time.NewTicker(interval)
	stopTicker := time.NewTicker(maxWait)

	success := false
	timeout := false

loop:
	for {
		select {
		case <-ticker.C:
			result, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err == nil {
				if result.Status == types.ReceiptStatusSuccessful {
					success = true
					ticker.Stop()
					stopTicker.Stop()
					break loop
				}
				if result.Status == types.ReceiptStatusFailed {
					ticker.Stop()
					stopTicker.Stop()
					break loop
				}
			}
		case <-stopTicker.C:
			ticker.Stop()
			stopTicker.Stop()
			timeout = true
			break loop
		}

	}

	if timeout {
		return false, errors.New("timeout")
	}
	return success, nil
}

func (bpc *BitspawnPoaClient) getAuth(privKeyHex string, client *ethclient.Client, gasLimit uint64) (*bind.TransactOpts, common.Address, error) {

	privateKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		bpc.log.Error(err)
		return nil, common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		bpc.log.Error("error casting public key to ECDSA")
		return nil, common.Address{}, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		bpc.log.Error(err)
		return nil, common.Address{}, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		bpc.log.Error(err)
		return nil, common.Address{}, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit   // in units
	auth.GasPrice = gasPrice   // big.NewInt(2000000)

	return auth, fromAddress, nil
}
