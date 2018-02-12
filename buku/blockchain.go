package buku

//Blockchain model
type Blockchain struct {
	Blocks []*Block
}

//AddBlock append new block to the chain
func (bc *Blockchain) AddBlock() {
	chainLength := len(bc.Blocks)
	previousBlock := bc.Blocks[chainLength-1]
	newBlock := CreateBlock(previousBlock.Hash, previousBlock.Index+1)

	if isValidNewBlock(previousBlock, newBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		//fail, error message
	}

}

func isValidNewBlock(prevBlock *Block, newBlock *Block) bool {

	if prevBlock.Index+1 != newBlock.Index {
		return false
	} else if prevBlock.Hash != newBlock.PreviousHash {
		return false
	} else {
		//more validation needed i.e Hash calc
	}

	return true
}

//NewBlockChain create new blockchain, initializing chain with genesis block
func NewBlockChain() *Blockchain {
	blockChain := &Blockchain{
		[]*Block{
			CreateBlock(string([]byte{}), 0),
		},
	}

	return blockChain
}
