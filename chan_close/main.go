package main

import "fmt"

func main() {
	// 管道关闭之后就只读不写
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	close(intChan)
	// intChan <- 3 // 关闭后继续写入就会出错

	// 管道的遍历只能使用for range，不能使用for...i<len(intChan)或者cap(intChan)这种，会出现奇怪情况
	// 如果管道关闭了，可以使用for range正常遍历
	// 如果管道没有关闭，遍历会出现deadlock
	// 遍历时，保证管道关闭（如果是多协程情况下，其实是要求协程任务先完成，再取出）
	for v := range intChan {
		fmt.Println(v)
	}
}
