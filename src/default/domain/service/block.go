package service

import (
	"default/domain/model"
	"time"
)

// CreateBlock creates a Block
func CreateBlock(data string, prevBlockHash []byte) *model.Block {
	block := &model.Block{Data: []byte(data), Hash: []byte{}, PrevBlockHash: prevBlockHash, CreatedAt: time.Now().Unix()}
	block.SetHash()
	return block
}
