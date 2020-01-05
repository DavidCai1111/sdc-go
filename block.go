package sdcgo

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Block is a block in a blockchain.
type Block struct {
	TimeStamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// Serialize serializes the block.
func (b *Block) Serialize() ([]byte, error) {
	return Serialize(b)
}

// NewBlock creates a new block.
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		TimeStamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
	}

	return block
}

// DeserializeBlock deserializes the given bytes into a block.
func DeserializeBlock(b []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(b))

	if err := decoder.Decode(&block); err != nil {
		return nil, err
	}

	return &block, nil
}
