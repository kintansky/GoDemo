package main

import "fmt"

// 空接口的channel
type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 10)

	allChan <- 10
	allChan <- "testStr"
	allChan <- Cat{Name: "cat1", Age: 1}

	// 如果要 获取第三个数据只能先退出前两个数据
	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("%T, %v\n", newCat, newCat)
	// 注意虽然打印出来是一个Cat类型，但实际上在编译过程中，程序并不知道是一个Cat类型，仍然认为是一个接口类型
	// fmt.Println(newCat.Name)	// 错误
	// 所以需要先做一次类型断言
	a := newCat.(Cat)
	fmt.Println(a.Name)
}
