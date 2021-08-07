package main

import (
	"fmt"
	"sync"
	"time"
)

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine(s string ,wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func main() {
	//並列化
	//goroutineのプログラムが終わらなくてもメインが終われば終わる

	//それを避けるために使用
	var wg sync.WaitGroup
	wg.Add(1)

	go goroutine("world",&wg)
	normal("hello")

	wg.Wait()
}