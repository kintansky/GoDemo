package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前CPU数量
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
	fmt.Println("num=", cpuNum)
}
