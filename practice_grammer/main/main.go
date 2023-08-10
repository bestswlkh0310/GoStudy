package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("asdadad")
	bc.AddBlock("hello")

	for _, _block := range bc.BlockChain {
		fmt.Printf("PreHash: %x\n", _block.PreBlockHash)
		fmt.Printf("Data: %s\n", _block.Data)
		fmt.Printf("Hash: %x\n", _block.Hash)
		pow := NewProofOfWork(_block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
