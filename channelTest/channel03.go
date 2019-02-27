package main

import (
	"fmt"
)

func receive(receiver chan<- string, str string) {
	receiver <- str
}

func send(sender <-chan string, receiver chan<- string) {
	str := <-sender
	receiver <- str
}

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//关闭channel
	close(ch)
	res := 0
	for v := range ch {
		res += v
	}
	fmt.Println(res)
}
