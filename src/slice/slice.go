package main

import "fmt"

func main(){
	var s []byte
	s = make([]byte,5,5)
	fmt.Println("len(s)=",len(s),"cap(s)=",cap(s))   //len(s)= 5 cap(s)= 5
	//上面的写法等同于下面的简洁写法
	s1 := make([]byte,5)
	fmt.Println("len(s1)=",len(s1),"cap(s1)=",cap(s1))   //len(s1)= 5 cap(s1)= 5

	b := []byte{'a','b','c','d'}
	fmt.Println(b[1:4])  //[98 99 100]

	//匿名结构体类型的切片
	str := []struct {name,department string}{}
	fmt.Println(str)

	//属性和零值
	str1 := []string{8:"go",2:"python","java","C","C++","Php"}
	//6个被明确的指定为string类型，三个被填充为默认的string类型
	fmt.Println("str1=",str1)   //str1= [  python java C C++ Php  go]
	fmt.Println(len(str1))   //len=9
	fmt.Println(cap(str1))   //cap=9

	var str2 []string
	//str2 = make([]string,10)  //len=10
	fmt.Println("str2=",str2) //[]
	fmt.Println("len(str2)=",len(str2))   //len=0

	arr1 := [...]string{"Go","Python","Java","C","C++","PHP"}
	slice1 := arr1[:4]
	fmt.Println(slice1)	//[Go Python Java C]
	fmt.Println("len(slice1)=",len(slice1),"cap(slice1)=",cap(slice1))	////len(slice1)= 4 cap(slice1)= 6

	//slice索引值超出了切片值当前的长度，就会发生panic  ==>> 容量超限
	slice1 = slice1[:cap(slice1)]
	//panic: runtime error: slice bounds out of range
	//slice1 = slice1[:cap(slice1)+1]   //容量超限
	fmt.Println("---->>>",slice1)	// [Go Python Java C C++ PHP]

	//slice1扩展  slice1的值发生改变，底层数组中最右边的两个元素值被新追加的值改变， C++,PHP --> Rubby,Erlang
	slice1 = append(slice1,"Rubby","Erlang")
	fmt.Println("append slice1:",slice1)	//[Go Python Java C Rubby Erlang]
	slice3 := append(slice1,"Rubby1","Erlang1")
	fmt.Println("append slice3:",slice3)	//[Go Python Java C C++ PHP Rubby Erlang Rubby1 Erlang1]


	slice2 := arr1[3:]
	fmt.Println(slice2)	//[C C++ PHP]
	fmt.Println("len(slice2)=",len(slice2),"cap(slice2)=",cap(slice2))	//len(slice2)= 3 cap(slice2)= 3
	//invalid slice index -2 (index must be non-negative)
	//slice2 = slice1[-2:]		//	切片的下标必须是正数
	fmt.Println("------>>>>>>>>",slice2)

}

