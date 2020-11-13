package main

import "fmt"

func a() {
	fmt.Println("func a")
}

// recover 必须和defer搭配使用
func b() {
	// 在panic异常退出之前注册defer指定一个匿名函数来尝试恢复
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b error")
			fmt.Println(err)
		}
	}()
	panic("func b") // panic 抛出
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()

}
