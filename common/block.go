package common

// Block model
type Block struct {
	Index        int64
	Transactions []Transaction
	PreviousHash string
	Hash         string
	Timestamp    int64
}
