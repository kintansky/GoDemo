package main

import "fmt"

func test() {
	// 错误处理
	defer func() {
		err := recover() // 使用recover捕获异常
		if err != nil {  // 判断err是否为空，不为空真的时有错
			fmt.Println(err)
			// 错误处理主题
		}
	}()
	num1 := 10
	num2 := 0
	// 这里发生panic
	res := num1 / num2
	fmt.Println(res)
}

func main() {
	test()
}
