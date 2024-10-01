package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	// 创建一个临时文件
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()

	// 写入一些内容到文件
	_, err = tmpFile.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}

	// 获取文件名
	fileName := tmpFile.Name()

	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 删除文件
	err = os.Remove(fileName)
	if err != nil {
		panic(err)
	}

	// 尝试读取文件内容
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("读取文件内容：%s\n", string(buf[:n]))

	// 等待一段时间
	time.Sleep(1 * time.Second)

	// 尝试再次读取文件内容
	n, err = file.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Printf("再次读取文件内容：%s\n", string(buf[:n]))
}
