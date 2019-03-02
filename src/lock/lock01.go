package main

import (
	"log"
	"time"
)

//多个gorouting，通过共享变量实现通信
var name = "Jone"

func printName() {
	log.Println("name is:", name)
}

func main() {
	for i := 0; i < 3; i++ {
		go printName()
	}
	time.Sleep(1 * time.Second)
	//修改名字
	name = "Tom"
	for i := 0; i < 3; i++ {
		go printName()
	}
	time.Sleep(1 * time.Second)
}
