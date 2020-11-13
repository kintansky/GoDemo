package main

import "fmt"

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1) // 10进制输出
	fmt.Printf("%b\n", i1) // 2进制输出
	fmt.Printf("%o\n", i1) // 8进制输出
	fmt.Printf("%x\n", i1) // 16进制输出
	i2 := 077              // 8进制以0开头
	fmt.Printf("%d\n", i2)
	i3 := 0x123456 // 16进制0x开头
	fmt.Printf("%d\n", i3)
	fmt.Printf("%T\n", i3) // 输出类型 int

	// 声明int8 类型
	var i4 int8 = 9 //或者i4 := int8(9)，否则简短声明会使用默认的int
	fmt.Printf("%T\n", i4)

}
