package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 直接Read读取
func readFile() {
	fileObj, err := os.Open("./a.txt")
	if err != nil {
		fmt.Printf("open failed, err: %v\n", err)
		return
	}
	defer fileObj.Close() // 关闭文件
	for {                 // 循环读取
		var tmp = make([]byte, 3) // 最长128字节的切片，相当于每次只读8字节
		n, err := fileObj.Read(tmp)
		if err == io.EOF { // 结束读取
			// 读取到最后如果还有不足128个字节的，也要取出来
			fmt.Printf("read %d bytes\n", n)
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("Read Failed, err: %v\n", err)
			return
		}
		fmt.Printf("read %d bytes\n", n)
		fmt.Println(string(tmp[:n])) // 因为tmp是一个byte类型，所以需要转换成string
	}
}

// 使用buf读取
func readFileByBufio() {
	fileObj, err := os.Open("./a.txt")
	if err != nil {
		fmt.Printf("open failed, err: %v\n", err)
		return
	}
	defer fileObj.Close() // 关闭文件

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(line)
			return
		}
		if err != nil {
			fmt.Printf("read file by buf failed, err: %v", err)
			return
		}
		fmt.Println(line)
	}
}

// 整个文件读取，大文件不推荐
func readFileByIoUtil() {
	content, err := ioutil.ReadFile("./a.txt")
	if err != nil {
		fmt.Printf("read file by ioutil, err: %v", err)
		return
	}
	fmt.Println(string(content))
}

// 写入
func write() {
	fileObj, err := os.OpenFile("./a.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// 注意标记位的区别
	if err != nil {
		fmt.Printf("Write error, err:%v", err)
		return
	}
	defer fileObj.Close()
	str := "tjx"
	fileObj.Write([]byte(str))
	fileObj.WriteString("hello")
}

// buf写入
func writeByBufio() {
	fileObj, err := os.OpenFile("./a.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// 注意标记位的区别
	if err != nil {
		fmt.Printf("Write error, err:%v", err)
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("buf write1")
	writer.WriteString("buf write2")
	writer.Flush() // 注意buf需要在最后调用flush，将缓存区内容写入文件
}

// write ioUitl
func writeByIoUntil() {
	str := "write by ioutil"
	err := ioutil.WriteFile("./a.txt", []byte(str), 0644) // 但是这里面就没有追加、清理等操作，会直接清空再写入
	if err != nil {
		fmt.Println("write file failed, err:%v\n", err)
		return
	}
}

func main() {
	// readFileByBufio()
	// readFileByIoUtil()
	// writeByBufio()
	writeByIoUntil()

}
