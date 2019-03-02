package main

import (
	"fmt"
	"time"
)

func strWorker(ch chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("get into strWorker...")
	ch <- "hello"
}

func intWorker(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("get into intWorker...")
	ch <- 10
}

func main() {
	chStr := make(chan string)
	chInt := make(chan int)
	go func() {
		strWorker(chStr)
	}()

	go func() {
		intWorker(chInt)
	}()

	for i := 0; i < 2; i++ {
		select {
		case <-chStr:
			fmt.Println("get value from chStr...")
		case <-chInt:
			fmt.Println("get value from chInt...")
		}
	}
}
