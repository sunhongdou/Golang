package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		number []int
		mux    sync.Mutex
	)

	for i := 0; i < 9; i++ {
		go func(v int) {
			//使用锁，因为多个gorouting对number存在竞争关系
			mux.Lock()
			number = append(number, v)
			mux.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("number is :", number)
}
