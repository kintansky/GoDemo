package main

import (
	"fmt"
	"sync"
)

// 如果使用全局变量进行不同协程之间的操作，注意需要加锁
// 如果使用channel也可以实现
var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}
	// 加入互斥锁，锁住全局资源
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i < 20; i++ {
		go test(i)
	}
	// 读取的时候也许要加锁
	lock.Lock()
	for i, r := range myMap {
		fmt.Printf("MAP[%d]: %d\n", i, r)
	}
	lock.Unlock()
}
