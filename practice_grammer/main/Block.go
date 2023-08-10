package main

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"strconv"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	PreBlockHash []byte
	Hash         []byte
	Nonce        int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) // 타임스탬프 -> 스트링 -> 바이트배열
	var headers []byte = bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

const targetBits = 16

type ProofOfWork struct {
	Block *Block

	target *big.Int
}
