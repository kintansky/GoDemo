package main

import "fmt"

func modify(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100
}

func main() {
	a := 10
	fmt.Printf("%p, %v\n", &a, a)     // %p 代表地址
	ptr := &a                         // ptr 存的是a的内存地址，当然是使用一个b自己的地址存放a的地址
	fmt.Printf("%p, %v\n", &ptr, ptr) // 第一个是ptr本来的地址，第二个是ptr存的值，就是a的地址
	fmt.Println(*ptr)

	modify(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}
