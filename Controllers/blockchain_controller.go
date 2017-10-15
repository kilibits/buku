package Controllers

import (
	"time"

	Models "../Models"
)

//Blocky alias to blockchain type
type Blocky Models.Blockchain

//NewBlockChain create new blockchain
func NewBlockChain() *Blocky {
	blockChain := &Blocky{
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
		Hash:         "TODO",
	}
	return block
}

//AddBlock append new block to the chain
func (bc *Blocky) AddBlock() {
	chainLength := len(bc.Blocks)
	previousBlock := bc.Blocks[chainLength-1]
	bc.Blocks = append(bc.Blocks, CreateBlock(previousBlock.Hash))

}
