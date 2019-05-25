package service

import (
	"default/domain/model"
)

// Blockchain is an array of Block
type Blockchain struct {
	Blocks []*model.Block
}

// CreateBlockchain creates a Blockchain
func CreateBlockchain() *Blockchain {
	return &Blockchain{[]*model.Block{CreateGenesisBlock()}}
}

// AddBlock is add Block in Blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlockHash := bc.Blocks[len(bc.Blocks)-1].Hash
	bc.Blocks = append(bc.Blocks, CreateBlock(data, prevBlockHash))
}
