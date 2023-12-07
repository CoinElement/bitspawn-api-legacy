/*

 */

package services

import "math/big"

func ConvertEthToWei(eth string) *big.Int {

	ethFloat, ok := Zero().SetString(eth)
	if !ok {
		return nil
	}

	exp, _ := Zero().SetString("1000000000000000000.0")

	weiFloat := Zero().Mul(ethFloat, exp)

	weiInt, _ := weiFloat.Int(nil)

	return weiInt
}

func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(512)
	return r
}
