package main

import (
	"fmt"
)

func main() {
	//案例一：非缓冲通道，需要至少两个gorouting来存入和读取，否则就会阻塞
	//	ch := make(chan string)
	//	go func() {
	//		ch <- "str"
	//	}()

	//	fmt.Println(<-ch)
	//案例二：缓冲通道，则不需要另外再起一个gorouting
	ch := make(chan string, 2)
	ch <- "str"
	fmt.Println(<-ch) //输出str
}
