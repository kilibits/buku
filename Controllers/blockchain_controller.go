package Controllers

import (
	"math/big"
	"time"

	Models "../Models"
)

//Chain alias to blockchain type
type Chain Models.Blockchain

const targetBits = 24

//NewBlockChain create new blockchain
func NewBlockChain() *Chain {
	blockChain := &Chain{
		[]*Models.Block{
			CreateBlock(string([]byte{})),
		},
	}

	return blockChain
}

//CreateBlock - create a new block
func CreateBlock(previousBlockHash string) *Models.Block {
	block := &Models.Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		PreviousHash: previousBlockHash,
		Transactions: []Models.Transaction{},
	}

	block.BlockHash()
	return block
}

//AddBlock append new block to the chain
func (bc *Chain) AddBlock() {
	chainLength := len(bc.Blocks)
	previousBlock := bc.Blocks[chainLength-1]
	bc.Blocks = append(bc.Blocks, CreateBlock(previousBlock.Hash))
}

func NewProofOfWork(block *Models.Block) *Models.Proof {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	proof := &Models.Proof{block, target}

	return proof
}
