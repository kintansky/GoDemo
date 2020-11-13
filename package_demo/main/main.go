package main

import (
	"fmt"
	// 导入包从GOPATH下的src之后开始, 可以添加js
	js "demo/package_demo/calc" // 注意所有平台下路径斜线都是 /
)

var test int = 100

func main() {
	fmt.Println("hello", js.Name) // 调用只需要使用包名.xxx 就可以调用，注意可见性即可
}

// 同样main中也可以定义init函数，也是在所有声明之后执行，外部导入的包的init会先于main的init执行
func init() {
	fmt.Println("main.init")
	fmt.Println(test)
	/* 执行结果
	calc.init // 外部包init
	tjx
	main.init // 当前包init
	100
	hello tjx
	*/
}
