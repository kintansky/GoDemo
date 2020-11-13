package main

import "fmt"

// chan 的作用是给多个goroutine之间通信
func main() {
	// 带缓冲区的通道：异步通道
	ch1 := make(chan int, 1) // 通道时引用类型，必须使用make初始化, 1为缓冲区大小
	ch1 <- 10                // 将值发送到chan
	x := <-ch1               // 从chan取值
	fmt.Println(x)
	// len(ch1)	// 求缓冲区元素的数量
	// cap(ch1)	// 求缓冲区的容量
	close(ch1)

	// 无缓冲区的通道：同步通道
	var ch2 chan int
	ch2 = make(chan int)
	// ch2 <- 10 // 无缓冲区的这样写会发生deadlock

	// 注意死锁情况
	// 1、如果写入超过容量，死锁，因此需要边写边取
	// 2、在没有使用协程的情况下，如果继续取值空的管道，死锁
}
