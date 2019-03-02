package main

import (
	"fmt"
	"time"
	"sync"
)

var(
	myMap = make(map[int]int,10)
	//申明一个全局的互斥锁
	lock sync.Mutex
)

func test(n int){
	res := 1
	for i:=1;i<=n;i++{
		res *= i
	}
	//写入时加锁
	lock.Lock()
	myMap[n] =res
	//写入完成解锁
	lock.Unlock()
}

func main(){
	//开启多个协程来完成这个任务
	for i:=1;i<=10;i++{
		//问题：多个协程往同一个map中写数据，需要做互斥操作
		go test(i)
	}

	//如果休眠十秒，等待协程完成
	time.Sleep(10*time.Second)

	//输出结果,发现没有输出数据，原有是主线程已经退出，协程还没有完成任务，只能强制结束
	//读的时候加锁
	lock.Lock()
	for i,v :=range myMap{
		fmt.Printf("map[%d] = %d",i,v)
	}
	//读完解锁
	lock.Unlock()
}
