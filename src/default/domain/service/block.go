package service

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"

	"default/domain/model"
)

// HashBlock is extends model.Block
type HashBlock model.Block

// CreateHash creates SHA-256 hash based on Block
func (b *HashBlock) CreateHash() {
	nonce := []byte(strconv.FormatInt(b.CreatedAt, 10))
	hash := sha256.Sum256(bytes.Join([][]byte{b.Data, b.PrevBlockHash, nonce}, []byte{}))
	b.Hash = hash[:]
}

// CreateBlock creates a Block
func CreateBlock(data string, prevBlockHash []byte) *HashBlock {
	block := &HashBlock{[]byte(data), []byte{}, prevBlockHash, time.Now().Unix()}
	block.CreateHash()
	return block
}
