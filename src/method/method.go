package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

//给结构体绑定一个方法
func (s Student) sayHello() {
	fmt.Println("hello")
}

//对于方法的接收者，既可以是值，也可以是指针，值类型对接受者是无法改变的，要想对接受者进行改变必须是指针类型，如*Student
func (s *Student) AddAge() {
	s.Age += 1
}
func (s Student) PrintAge() {
	fmt.Println(s.Name, "age is", s.Age)
}

func main() {
	s := Student{
		"Johe",
		8,
	}
	s.PrintAge()
	s.AddAge()
	s.PrintAge()
	//s.sayHello()
}
