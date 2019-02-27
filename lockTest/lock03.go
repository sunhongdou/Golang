package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
			stat = map[string]int{
			"a": 1,
		}
		//定义同步锁
		mux   sync.Mutex
		muOps uint64

		stat2 = map[string]int{
			"a": 1,
		}
		//定义读写锁
		rw    sync.RWMutex
		rwOps uint64
	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				mux.Lock()
				_ = stat["a"]
				mux.Unlock()
				atomic.AddUint64(&muOps, 1)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				rw.RLock()
				_ = stat2["a"]
				rw.RUnlock()
				atomic.AddUint64(&rwOps, 1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("muOps is", muOps)
	fmt.Println("rwOps is ", rwOps)
}
