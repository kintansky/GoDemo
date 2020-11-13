package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d, start job %d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker %d, finish job %d\n", id, job)
	}
}

// 限制线程数
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 10)

	// 开启3个goroutine
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results) // 开启goroutine的时候jobs和results都是空的
	}

	// 发送5个任务
	jobNum := 5
	for i := 0; i < jobNum; i++ {
		jobs <- i // 每发送一次值，3个worker中接收到的值的worker就会工作
	}
	close(jobs)
	// 最后把值worker结果取出来
	// 但因为result 还没有close，这里如果使用range取值，会一直遍历到空的位置也不停，最后deadlock
	// for ret := range results {
	// 	fmt.Println(ret)
	// }
	// 在没有使用协程的情况下，如果继续取值空的管道，死锁
	// 所以只能强制限制取值的个数为jobNum
	for i := 0; i < jobNum; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
