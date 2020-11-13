package main

import "fmt"

// 运算符
func main() {
	// 位运算
	a := 1 // 对应二进制001
	b := 5 // 对应二进制101
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)  // 异或
	fmt.Println(1 << 2) // 左移两位 成为二进制100，即4
	fmt.Println(4 >> 2) // 右移两位 成为二级制1，即1

	a1 := 2
	b1 := 4.0
	fmt.Println(float64(a1) / b1)

}
