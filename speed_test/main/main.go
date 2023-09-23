package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1) // WaitGroup 카운터 증가
		go gogo(&wg)
	} // 10조 번의 연산
	wg.Wait()
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("실행 시간: %v\n", executionTime)
}

func gogo(wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for i := 0; i < 100_0000_0000; i++ {
		sum++
	}
	fmt.Println(sum)
}
