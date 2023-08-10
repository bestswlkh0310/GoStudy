package main

import (
	"fmt"
	"strconv"
)

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) Run() {

	for {
		var mode string
		fmt.Println("print - 0 / add - 1")
		fmt.Scanln(&mode)
		if mode == "0" {
			cli.printChain()
		} else if mode == "1" {
			var data string
			fmt.Println("data를 입력해주세요")
			fmt.Scanln(&data)
			cli.addBlock(data)
		} else {
			fmt.Println("제대로 입력하셈")
		}
	}
}

func (cli *CLI) printUsage() {
	fmt.Println("ㅂㅇㅂㅇ")
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()
		fmt.Printf("이전 해쉬: %x\n", block.PreBlockHash)
		fmt.Printf("의 데이터: %s\n", block.Data)
		fmt.Printf("지금 해쉬: %x\n", block.Hash)

		pow := NewProofOfWork(block)
		fmt.Printf("작업 인증!: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PreBlockHash) == 0 {
			break
		}
	}
}
