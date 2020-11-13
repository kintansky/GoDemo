package main

// 解决chan不关闭情况下的遍历问题
import (
	"fmt"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 10)
	for i := 0; i < 5; i++ {
		stringChan <- fmt.Sprintf("hello%d", i)
	}

	// 传统方法在遍历管道时要求管道先关闭，
	// 但也不排除有些情景下，无法确定关闭管道的实际

	for {
		// 通过下面这种办法，如果intChan不关闭，不会一直导致取出最后一个时出现deadlock，而会到下一个case
		select {
		case v := <-intChan:
			fmt.Printf("从intChan取出的数据为%v\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan取出的数据为%v\n", v)
		default:
			fmt.Printf("都没有\n")
			return
		}
	}
}
