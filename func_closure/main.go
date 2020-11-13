package main

// 闭包
import (
	"fmt"
	"strings"
)

// 闭包：函数+引用环境
// 定义一个函数，返回值是函数
func a(name string) func() {
	s := "tjx"
	return func() {
		fmt.Println("hello", name, s) // 引用外层变量name
	}
}

// 文件后缀检测
func makeSuffixFunc(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}

// 多函数返回值
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		return base + i
	}
	sub := func(i int) int {
		return base - i
	}
	return add, sub
}

func main() {
	r := a("test") // r 包含一个匿名函数和一个外层变量
	r()            // 相当于直接执行了a函数内部的匿名函数

	// 文件后缀检测
	r2 := makeSuffixFunc(".txt") // 第一层传外层环境变量
	newName := r2("nihao")       // 第二层调用传内部函数的参数
	fmt.Println(newName)

	// 多函数返回值
	add, sub := calc(1000)
	ret1 := add(100)
	ret2 := sub(200)
	fmt.Println(ret1, ret2)
}
