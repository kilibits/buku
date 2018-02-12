package buku

import (
	"crypto/sha256"
	"strings"
	"time"
)

// Block model
type Block struct {
	Index        int64
	Transactions []Transaction
	PreviousHash string
	Hash         string
	Timestamp    int64
}

//BlockHash calculate SHA-256 hash for the block
func (b *Block) BlockHash() {
	headers := strings.Join([]string{b.PreviousHash, string(b.Timestamp)}, "")
	hash := sha256.Sum256([]byte(headers))
	b.Hash = string(hash[:])
}

//CreateBlock - create a new block
func CreateBlock(previousBlockHash string, index int64) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		PreviousHash: previousBlockHash,
		Transactions: []Transaction{},
	}

	block.BlockHash()
	return block
}
