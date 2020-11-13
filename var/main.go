package main

import "fmt"

var h string = "tt" // 全局变量可以声明后不使用，非全局变量声明后必须使用
var (
	s1   string
	age  int
	isOk bool
)

func main() {
	s1 = "test"
	age = 20
	isOk = true
	fmt.Print(isOk)           // 不换行
	fmt.Printf("name:%s", s1) // 格式化输出，也不换行
	fmt.Println(age)          //换行

	// 类型推导
	var s2 = "hhh"
	fmt.Println(s2)

	// 简短变量声明，声明并赋值
	s3 := "哈哈哈" // 只能使用在函数内部，不能用于全局变量的声明
	fmt.Println(s3)

	// _匿名变量，表示忽略值，不占内存空间

}
