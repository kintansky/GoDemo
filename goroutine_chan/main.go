package main

// goroutine和chan的结合，利用chan进行协程之间的通讯，同时利用chan阻塞主线程
import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		// time.Sleep(time.Second * 2) // 模拟处理时间
		intChan <- i
		fmt.Println("writeData: ", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData: ", v)
		time.Sleep(time.Second * 1) // 模拟处理时间
	}
	// readData 任务完成，退出
	exitChan <- true
	close(exitChan)
}

// 经典的生产者消费者模型
func main() {
	intChan := make(chan int, 10) // 一个小管道处理大数据量，例如管道容量10，数据量50
	exitChan := make(chan bool, 1)

	// 并发读写
	go writeData(intChan)
	go readData(intChan, exitChan) // 如果没有消费者readData，因为intChan的容量是10，数据量比容量大，这时候就会在写入超过intChan长度的时候死锁

	// 利用exitChan阻塞主线程，不让主线程退出
	for ex := range exitChan {
		if ex {
			break
		}
	}
	// // 或者
	// for {
	// 	_, ok := <-exitChan
	// 	if !ok {
	// 		break
	// 	}
	// }

}
