package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 全局变量

func hello(i int) {
	// 将任务包装成函数，实现并发
	time.Sleep(2 * time.Second)
	fmt.Println("hello...", i)

	wg.Done() // 通知任务计数器-1
}

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)   // 任务计数牌+1
		go hello(i) // 单独开启一个goroutine执行hello
	}
	fmt.Println("hello main")
	wg.Wait() // 阻塞直至计数器归0再退出

	// 注意如果goroutine包含匿名函数闭包的时候，要注意传参的问题
	for i := 0; i < 100; i++ {
		wg.Add(1)
		// 匿名函数中的i因为是从外部取参，是一个闭包，所以如果不显性传参会导致参数出错，
		// go func() {
		// 	fmt.Println("hello...", i)
		// 	wg.Done()
		// }()
		// 正确写法，显性传参i
		go func(i int) {
			fmt.Println("hello...", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
