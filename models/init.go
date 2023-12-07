/*

 */

package models

import "math/big"

var PrizeAllocation map[int64][]*big.Int

func init() {

	PrizeAllocation = map[int64][]*big.Int{
		3: {
			big.NewInt(50),
			big.NewInt(30),
			big.NewInt(20),
		},
	}
}
