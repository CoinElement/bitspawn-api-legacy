/*

 */

package poa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/bitspawngg/bitspawn-api/contracts/pool"
	"github.com/bitspawngg/bitspawn-api/services"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (bpc *BitspawnPoaClient) GiftEth(userAddress string) (*types.Transaction, error) {
	client, err := bpc.getClient()
	if err != nil {
		return nil, err
	}

	privateKey, _ := crypto.HexToECDSA(bpc.faucetKey)

	auth, _, err := bpc.getAuth(bpc.faucetKey, client, services.DefaultGasLimit)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	value := big.NewInt(1000000000000000000) // 1 Eth in wei
	//func NewTransaction(nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *Transaction
	tx := types.NewTransaction(auth.Nonce.Uint64(), common.HexToAddress(userAddress), value, auth.GasLimit, auth.GasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	//ok, err := WaitTxSuccess(client, signedTx, time.Second*20, time.Second)
	//if !ok {
	//	return tx, errors.New("transaction failed")
	//}
	//if err != nil {
	//	return tx, errors.New("tx receipt timed out")
	//}
	return tx, nil
}

type ResponseBodyWalletAPI struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (bpc *BitspawnPoaClient) GiftSpwn(xAuthToken string) error {
	var responseBody ResponseBodyWalletAPI

	client := &http.Client{}
	req, _ := http.NewRequest("POST", bpc.walletUrl+"/freetoken", nil)
	req.Header.Add("X-Auth-Token", xAuthToken)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error calling wallet API: %w", err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &responseBody)
	if err != nil {
		return fmt.Errorf("Error in unmarshalling json payload from Wallet API: %w", err)
	}
	if !responseBody.Ok {
		return fmt.Errorf("Wallet API returns error: %s", responseBody.Message)
	}
	return nil
}

func (bpc *BitspawnPoaClient) Fund(privKeyHex string, address common.Address, amountInWei *big.Int) (*types.Transaction, error) {
	client, err := bpc.getClient()
	if err != nil {
		return nil, err
	}

	auth, _, err := bpc.getAuth(privKeyHex, client, services.DefaultGasLimit)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	bitspawn, err := pool.NewBitspawn(bpc.spwnContract, client)
	if err != nil {
		bpc.log.Error(err)
		return nil, err
	}

	tx, err := bitspawn.Transfer(auth, address, amountInWei)
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
