package main

import (
	"fmt"
)

func worker(ch chan string) {
	fmt.Println("get into worker ...")
	ch <- "str"
}

func main() {
	ch := make(chan string)
	go worker(ch)
	//channel阻塞，如果没有下面这行代码，worker函数就不会执行，因为main函数比gorouting快
	<-ch
}
