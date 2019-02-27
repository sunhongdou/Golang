package main

import (
	"fmt"
	"math"
)

//定义一个接口为匿名成员
type tmp interface {
	area() float32
}

//定义一个接口，有两个方法，求面积和周长
type geometry interface {
	//area() float32
	tmp
	perim() float32
}

//定义长方形结构体
type rect struct {
	len, width float32
}

//定义圆的结构体
type circle struct {
	radius float32
}

//定义长方形求面积的方法
func (r *rect) area() float32 {
	return r.len * r.width
}

//定义一个长方形求周长的方法
func (r *rect) perim() float32 {
	return (r.len + r.width) * 2
}

//求圆面积
func (c *circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

//求圆周长
func (c *circle) perim() float32 {
	return math.Pi * c.radius * 2
}

func show(name string, param interface{}) {
	switch param.(type) {
	case geometry:
		fmt.Printf("area of %v is %v\n", name, param.(geometry).area())
		fmt.Printf("perim of %v is %v\n", name, param.(geometry).perim())
	default:
		fmt.Println("wrong type")
	}

}
func main() {
	//定义一个长方体
	rect := &rect{
		len:   1,
		width: 2,
	}
	show("rect", rect)
	//定义一个圆形
	cir := &circle{
		radius: 2,
	}
	show("circle", cir)
}
