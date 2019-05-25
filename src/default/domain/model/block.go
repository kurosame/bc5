package model

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

// Block is transaction
type Block struct {
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
	CreatedAt     int64
}

// SetHash creates SHA-256 hash based on Block
func (b *Block) SetHash() {
	nonce := []byte(strconv.FormatInt(b.CreatedAt, 10))
	hash := sha256.Sum256(bytes.Join([][]byte{b.Data, b.PrevBlockHash, nonce}, []byte{}))
	b.Hash = hash[:]
}
