package service

import (
	"default/domain/model"
)

// Blockchain is an array of Block
type Blockchain struct {
	Blocks []*model.Block
}

// AddBlock is add Block in Blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlockHash := bc.Blocks[len(bc.Blocks)-1].Hash
	bc.Blocks = append(bc.Blocks, CreateBlock(data, prevBlockHash))
}
