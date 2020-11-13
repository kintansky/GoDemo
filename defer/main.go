package main

import "fmt"

func test(i int) int {
	defer fmt.Println("close", i)
	defer fmt.Println("close*", i)
	fmt.Println(i)
	return i * i
}

func main() {
	// defer 会在执行完后，按照倒序执行
	// fmt.Println("start")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("end")
	// defer 语句一般用于资源的释放、文件关闭、解锁、记录时间等

	for i := 0; i < 10; i++ {
		r := test(i)
		fmt.Println(r)
	}

}
