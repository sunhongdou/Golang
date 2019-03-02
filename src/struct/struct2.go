package main

import (
	"fmt"
)

type Person struct {
	Age  int
	Name string
}

//结构体内嵌
type Student struct {
	Person
}

func main() {
	stu := Student{
		Person{
			Age:  17,
			Name: "Jone",
		},
	}
	fmt.Println(stu)
	fmt.Println(stu.Name)
	fmt.Println(stu.Age)
}
