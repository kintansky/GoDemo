package main

// 利用并行找素数
import "fmt"

func putNum(intChan chan int) {
	for i := 0; i < 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok { // 管道读取完退出
			break
		}
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 { // 判断不是素数，flag标记，退出
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num // 是素数，加入结果中
		}
	}
	// primeChan不能再这里关闭，因为不知道别的协程是否完成
	// exitChan也不能在这里close，因为我们起了4个线程进行处理，所以不知道别的线程是否完成
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) // 保存结果的chan
	exitChan := make(chan bool, 4)    // 4个worker协程处理完成的标记chan

	go putNum(intChan)
	for i := 0; i < 4; i++ { // 设置4个协程来处理这里问题，处理结果各自放入primeChan，各自完成的结果放入exitChan
		go primeNum(intChan, primeChan, exitChan)
	}
	// // 最后阻塞主线程，并关闭primeChan
	// for i := 0; i < 4; i++ {
	// 	<-exitChan //因为如果别的线程没有完成，会阻塞再这里，所以不需要计数，直接扔掉即可
	// }
	// close(primeChan)
	go func() { // 因为chan在协程里面也是阻塞的，所以最好作为一个匿名函数操作
		for i := 0; i < 4; i++ {
			<-exitChan //因为如果别的线程没有完成，会阻塞再这里，所以不需要计数，直接扔掉即可
		}
		close(primeChan)
	}()

	// 遍历结果
	for v := range primeChan {
		fmt.Println(v)
	}
	// 或者
	// for {
	// 	v, ok := <-primeChan
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }
}
