package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"monster_name"`
	Age int `json:"monster_age"`
}
//序列化结构体
func MarshalStruct(){
	monster := Monster{Name:"tom",Age:14}
	//将结构体进行序列化
	data,err := json.Marshal(&monster)
	if err!=nil{
		fmt.Printf("marshal failed，err is %v",err)
	}
	//序列化后的结果
	fmt.Printf("monster marshal result=[%v]",string(data))
}

//序列化map
func MarshalMap(){
	var student map[string]interface{}
	student = make(map[string]interface{})
	student["name"] = "zhangsan"
	student["age"] = 18

	//将map进行序列化
	data,err := json.Marshal(student)
	if err!=nil{
		fmt.Printf("marshal failed，err is %v",err)
	}
	//序列化后的结果
	fmt.Printf("student marshal result=[%v]",string(data))
}

func MarshalSlice(){
	var slice []map[string]interface{}
	var map1 map[string]interface{}
	map1 = make(map[string]interface{})
	map1["name"] = "zhangsan"
	map1["age"] = 32

	slice =append(slice,map1)

	var map2 map[string]interface{}
	map2 = make(map[string]interface{})
	map2["name"] = "zhangsan"
	map2["age"] = 32
	map2["addr"] = []string{"beijing","shanghai"}

	slice =append(slice,map2)

	data,err := json.Marshal(slice)
	if err!=nil{
		fmt.Printf("marshal failed，err is %v",err)
	}
	//序列化后的结果
	fmt.Printf("student marshal result=[%v]",string(data))
}

func UnmarshalStruct(){
	str := "{\"monster_name\":\"tom\",\"monster_age\":14}"
	var monster Monster
	err := json.Unmarshal([]byte(str),&monster)
	if err!=nil{
		fmt.Printf("unmarshal failed,err is : [%v]",err)
	}
	fmt.Printf("monster name is : %v",monster.Name)
}

func main(){
	//map slice struct 序列化为json
	//MarshalStruct()

	//MarshalMap()

	//MarshalSlice()

	UnmarshalStruct()
}
