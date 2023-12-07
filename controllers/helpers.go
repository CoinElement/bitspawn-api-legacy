/*

 */

package controllers

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

func ConvertWeiToEth(wei string) *big.Float {

	weiFloat, ok := Zero().SetString(wei)
	if !ok {
		return nil
	}

	exp, _ := Zero().SetString("1000000000000000000.0")

	ethFloat := Zero().Quo(weiFloat, exp)

	return ethFloat
}

func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(512)
	return r
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
