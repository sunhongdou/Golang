package demo01

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func main() {
	ReadDirAll("../")
}

//遍历目录
func ReadDirAll(dirName string) {
	fileInfos, _ := ioutil.ReadDir(dirName)
	count := strings.Count(dirName, "/")
	for i, v := range fileInfos {
		temp := strings.Repeat("-", count)
		fmt.Printf("%s %d,%s(%t) \n", temp, i+1, v.Name(), v.IsDir())
		if v.IsDir() {
			ReadDirAll(path.Join(dirName, v.Name())) //递归调用
		}
	}
}
