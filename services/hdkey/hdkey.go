/*

 */

package hdkey

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/crypto"
	uuid "github.com/satori/go.uuid"
)

func GeneratePrivateKeyFromUUID(uuidString string) (string, error) {

	uuid, err := uuid.FromString(uuidString)
	if err != nil {
		return "", errors.New("bad token")
	}

	master := "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jP" +
		"PqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi"

	key, _ := hdkeychain.NewKeyFromString(master)

	for i := 0; i < 4; i++ {
		var value uint32
		value |= uint32(uuid[4*i])
		value |= uint32(uuid[4*i+1]) << 8
		value |= uint32(uuid[4*i+2]) << 16
		value |= uint32(uuid[4*i+3]) << 24

		key, err = key.Child(value)
		if err != nil {
			return "", errors.New("bad token")
		}
	}
	btcecKey, err := key.ECPrivKey()
	if err != nil {
		return "", errors.New("bad token")
	}

	return hex.EncodeToString(btcecKey.Serialize()), nil
}

func GeneratePublicAddressFromUUID(uuidString string) (string, error) {
	privateKey, err := GeneratePrivateKeyFromUUID(uuidString)
	if err != nil {
		return "", err
	}

	return GeneratePublicAddressFromPrivateKey(privateKey)
}

func GeneratePublicAddressFromPrivateKey(privateKey string) (string, error) {
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	publicKeyECDSA, ok := privKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}
	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return publicAddress, nil
}
