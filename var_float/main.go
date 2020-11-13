package main

import "fmt"

func main() {
	// float 只用float32或者float64 没有float类型。默认使用的是float64
	f1 := 1.12345
	fmt.Printf("%Tvalue:%v", f1, f1)
	// float32 和float64 是两种类型，不能直接互相赋值

}
