package Models

import "math/big"

//Proof model
type Proof struct {
	Block  *Block
	Target *big.Int
}
