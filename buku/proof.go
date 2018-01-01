package buku

import "math/big"

//Proof model
type Proof struct {
	Block  *Block
	Target *big.Int
}

func (proof *Proof) PrepareData(nonce int) string {
	// data := bytes.Join(
	// 	[][]byte{
	// 		proof.Block.PreviousHash,
	// 	}
	// )

	return "hello"
}

const targetBits = 24

func NewProofOfWork(block *Block) *Proof {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	proof := &Proof{block, target}

	return proof
}
