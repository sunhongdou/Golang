package demo10

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skll string
}

//给monster绑定一个Store方法，可以将monster对象，序列化后保存到文件中
func (this *Monster) Store() bool{
	data,err := json.Marshal(this)
	if err!=nil{
		fmt.Println("marshal failed,err is : [%v]",err)
		return  false
	}
	filePath := "d:/monster.ser"
	err = ioutil.WriteFile(filePath,data,0666)
	if err !=nil {
		fmt.Println("write file failed,err is : [%v]",err)
		return false
	}
	return  true
}

//反序列化
func (this *Monster) Restore() bool {
	filePath := "d:/monster.ser"
	data,err := ioutil.ReadFile(filePath)
	if err!=nil{
		fmt.Println("read file failed,err is : [%v]",err)
		return  false
	}

	err = json.Unmarshal(data,this)
	if err!=nil{
		fmt.Println("unmarshal failed,err is : [%v]",err)
		return  false
	}
	return  true
}
