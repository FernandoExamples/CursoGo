package main

import (
	"fmt"
	"golang-blockchain/models"
	"strconv"
)

func main() {
	chain := models.InitBlockChain()

	chain.AddBlock("First block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("Third block after genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := models.NewProof(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
