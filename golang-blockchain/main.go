package main

import (
	"fmt"
	"golang-blockchain/models"
)

func main() {
	chain := models.InitBlockChain()

	chain.AddBlock("First block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("Third block after genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
