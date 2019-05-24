package service

import (
	"bytes"
	"crypto/sha256"
	"strconv"

	"default/domain/model"
)

// HashBlock is extends model.Block
type HashBlock model.Block

// CreateHash create SHA-256 hash based on Block
func (b *HashBlock) CreateHash() {
	nonce := []byte(strconv.FormatInt(b.CreatedAt, 10))
	hash := sha256.Sum256(bytes.Join([][]byte{b.PrevBlockHash, b.Data, nonce}, []byte{}))
	b.Hash = hash[:]
}
