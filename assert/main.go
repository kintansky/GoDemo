package main

import "fmt"

type B struct {
}

func main() {
	var a interface{}
	b := B{}
	a = b // 把结构体赋值给空接口可以，但是无法直接把接口类型赋值给具体类型
	fmt.Println(a)

	var b2 B
	// b2 = a // 错，无法直接把接口赋值给类型，需要断言
	b2 = a.(B) // 类型断言，把接口类型断言成类型B
	fmt.Printf("%v, %T\n", b2, b2)

	// var x interface{}
	// var b3 float32 = 1.1
	// x = b3
	// y := x.(float32) // 注意，断言时必须与原类型匹配，如果转换成float64会panic
	// fmt.Printf("%v, %T\n", y, y)

	// 带检测类型断言
	var x interface{}
	var b3 float32 = 1.1
	x = b3
	y, ok := x.(float64) // 注意，断言时必须与原类型匹配，如果转换成float64会panic
	if ok {              // 加入检测
		fmt.Printf("%v, %T\n", y, y)
	} else {
		fmt.Println("assert failed")
	}

}
