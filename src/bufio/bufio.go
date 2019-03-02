package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//构建带有缓存的Reader对象：bufio.Reader
func testReader() {
	fileName := "../.idea/Test.iml"
	file1, _ := os.Open(fileName)
	reader1 := bufio.NewReader(file1)

	for {
		s1, err := reader1.ReadString('\n')
		fmt.Print(s1)
		if err == io.EOF {
			fmt.Println("读取完毕！")
			break
		}
	}

	file1.Close()
}

func testWriter() {
	fileName2 := "../.idea/Test.iml"
	file2, _ := os.Open(fileName2)
	reader2 := bufio.NewReader(file2)

	fileName3 := "../demo02/test.iml"
	file3, _ := os.OpenFile(fileName3, os.O_WRONLY|os.O_CREATE, os.ModePerm)

	writer1 := bufio.NewWriter(file3)

	for {
		bs, err := reader2.ReadBytes(' ')
		writer1.Write(bs)
		writer1.Flush()
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
	}
	file2.Close()
	file3.Close()
}

func main() {
	//testReader()
	testWriter()
}
