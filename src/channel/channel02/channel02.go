package main

import (
	"fmt"
)

//单项chanel
func receive(receiver chan<- string, str string) {
	receiver <- str
}

func send(sender <-chan string, receiver chan<- string) {
	str := <-sender
	receiver <- str
}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	receive(ch1, "hello") //将hello传递给ch1
	send(ch1, ch2)        //将ch1中的值传递给ch2
	fmt.Println(<-ch2)
}
