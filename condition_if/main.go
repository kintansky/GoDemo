package main

import "fmt"

func main() {
	// if
	var score = 90
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// if体内可以像C一样声明，但是注意作用域只在if内
	if score := 60; score > 90 {
		fmt.Println("A")
	} else {
		fmt.Println("C")
	}

	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 可以省略初始语句
	i := 10
	for ; i > 0; i-- {
		fmt.Println(i)
	}
	// 省略自增语句
	i = 10
	for j := 10; i > 0; {
		fmt.Println(i, j)
		i--
	}
	// 省略初始化语句和自增语句
	i = 10
	for i > 0 {
		fmt.Println(i)
		i--
	}

	// break
	for n := 0; n < 10; n++ {
		if n == 3 {
			break
		} else {
			fmt.Println(n)
		}
	}
	//continue
	for n := 0; n < 10; n++ {
		if n < 3 {
			continue
		} else {
			fmt.Println(n)
		}
	}
}
