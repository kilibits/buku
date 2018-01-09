package buku

import (
	"math/big"
)

//Proof model
type Proof struct {
	Block  *Block
	Target *big.Int
}

const targetBits = 24

func (proof *Proof) prepareData(nonce int) string {
	return ""
}

//NewProofOfWork -
func NewProofOfWork(block *Block) *Proof {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	proof := &Proof{block, target}

	return proof
}
