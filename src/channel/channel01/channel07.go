package main

import (
	"fmt"
	"time"
)

func sayHello(){
	for i:=0;i<10;i++{
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test(){
	//防止该协程发生异常，需要加defer和recover()
	defer func(){
		//捕获test抛出的panic,这样就不会影响其他协程的正常执行
		if err :=  recover();err!=nil{
			//test occur panic,err is : assignment to entry in nil map
			fmt.Println("test occur panic,err is :",err)
		}
	}()
	//定义一个map
	var myMap map[int]string
	//assignment to entry in nil map
	myMap[0] = "golang"   //用法错误，map需要先make
}

func main(){
	//实际开发中,即使一个协程发生panic，不希望影响到另外一个协程，这里可以使用recover机制解决。
	go sayHello()
	go test()

	for i:=0;i<10;i++{
		fmt.Printf("main(),ok = %d\n",i)
		time.Sleep(time.Second)
	}
}
