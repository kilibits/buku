package main

import (
	"fmt"
	"time"

	. "./buku"
)

func main() {
	blockChain := NewBlockChain()
	blockChain.AddBlock()
	blockChain.AddBlock()
	blockChain.AddBlock()

	for _, block := range blockChain.Blocks {
		fmt.Println("Block Information")
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0))
		fmt.Println()
	}
}
