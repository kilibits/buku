package buku

//Blockchain model
type Blockchain struct {
	Blocks []*Block
}

//AddBlock append new block to the chain
func (bc *Blockchain) AddBlock() {
	chainLength := len(bc.Blocks)
	previousBlock := bc.Blocks[chainLength-1]
	bc.Blocks = append(bc.Blocks, CreateBlock(previousBlock.Hash))
}

//NewBlockChain create new blockchain
func NewBlockChain() *Blockchain {
	blockChain := &Blockchain{
		[]*Block{
			CreateBlock(string([]byte{})),
		},
	}

	return blockChain
}
