package main

import "fmt"

/*
两个goroutine，两个channel
1、生成0-100数字发给ch1
2、从ch1取出数据计算平方，把结果发送到ch2
*/

func f1(ch chan<- int) { // 单向通道 一般用于函数传参的时候，这样规定了以后，在函数体内只能往通道发，不能从通道取值
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch) // 关闭后不能在操作，但能取值出来，直到返回false，见f1操作
}

func f2(ch1 <-chan int, ch2 chan<- int) { // 单向通道，ch1只可取值不可接收；ch2只能接收不可取值
	// 从通道中取值的方式1,死循环遍历
	for {
		tmp, ok := <-ch1 // 可以一直取值直到取完为止
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)
	go f1(ch1)
	go f2(ch1, ch2)

	// 从通道取值的方法2，如果ch2没有close不能通过以下方法取出值，否则死锁
	for ret := range ch2 {
		fmt.Println(ret)
	}

}
