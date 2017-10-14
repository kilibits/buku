package Controllers

import (
	"time"

	Models "../Models"
)

//CreateBlock - create a new block
func CreateBlock() *Models.Block {
	block := &Models.Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		PreviousHash: "TODO",
		Hash:         "TODO",
	}
	return block
}
