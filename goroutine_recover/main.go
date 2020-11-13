package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("hello")
	}
}

func test() {
	// 如果其中一个协程有错会影响整个线程执行
	// 采用recover捕获即可，既可以不影响其他所有线程
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test func error", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "test"
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}
