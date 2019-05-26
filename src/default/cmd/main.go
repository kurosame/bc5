package main

import (
	"default/domain/service"
	"fmt"
)

func main() {
	bc := service.CreateBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, b := range bc.Blocks {
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("Prev. hash: %x\n", b.PrevBlockHash)
		fmt.Println()
	}
}
