package main

type Chain struct {
	BlockChain []*Block
}

func (bc *Chain) AddBlock(data string) {
	preBlock := bc.BlockChain[len(bc.BlockChain)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.BlockChain = append(bc.BlockChain, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte{})
}

func NewBlockChain() *Chain {
	return &Chain{[]*Block{NewGenesisBlock()}}
}
