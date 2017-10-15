package Controllers

import (
	"time"

	Models "../Models"
)

//Chain alias to blockchain type
type Chain Models.Blockchain

//Block alias to block type
type Block Models.Block

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
