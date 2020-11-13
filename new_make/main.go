package main

import "fmt"

func main() {
	/*
		// 错误写法, 引用类型直接声明的时候是不会分配内存地址的，是一个nil，nil是无法赋值
		var a *int	// nil
		*a = 100
		fmt.Println(*a)
	*/
	var a *int
	var b int = 100
	a = &b
	fmt.Println(*a, a, &b)
	// 使用new函数，new函数返回对应类型的指针，该指针对应的值为该类型的零值
	c := new(int) // 使用new初始化
	// // 等效于
	// var c *int
	// c = new(int)
	fmt.Printf("%T\n", c) // *int

	// make 只能用于slice， map，chan的内存分配，而且返回的是类型本身不是指针
	var d map[string]int         // 引用类型只声明还未初始化
	d = make(map[string]int, 10) // 初始化
	d["test"] = 100              //如果不初始化，会导致nil赋值，panic
	fmt.Println(d)

}
