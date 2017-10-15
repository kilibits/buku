package main

import (
	Controllers "./Controllers"
)

func main() {
	blockChain := Controllers.NewBlockChain()
	blockChain.AddBlock()
}
