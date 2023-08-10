package main

import (
	"sync"
	"time"
)

type Person struct {
	Name string
	Age  int
}

var mu = sync.Mutex{}

type Movable interface {
	Run()
}

func main() {
	person1 := Person{"h", 1}
	print(person1.Age)
	print("\n")
	defer func() {
		println("Hello")
	}()
	print(Add(1, 3))
	go func() { PrintSOS() }()
	go func() { PrintSOS() }()
	time.Sleep(time.Second * 4)
}

func PrintSOS() {
	//mu.Lock()
	//defer mu.Unlock()
	println("Help")
	time.Sleep(time.Second * 2)
}

func Add(a int, b int) (result int) {
	result = a + b
	return
}
