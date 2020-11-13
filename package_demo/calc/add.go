package calc

import "fmt"

// Name 公开变量
var Name = "tjx"

func add(x, y int) int {
	return x + y
}

// Sub 同一个文件夹下的.go不需要做额外导入即可使用其他的.go文件
func Sub(x, y int) int {
	return sub(x, y)
}

// 包导入时自动执行，执行时机为当前包所有声明之后
func init() {
	fmt.Println("calc.init")
	fmt.Println(Name)
}
