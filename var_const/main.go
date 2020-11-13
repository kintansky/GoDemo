package main

import "fmt"

const pi = 3.1415 // 常量声明后不能修改
const (
	OK       = 200
	NOTFOUND = 404
)
const (
	n1 = 100
	n2 // 如果常量批量声明中忽略了值，表示与上面一个常量一致
	n3
)

// iota 常量计数器，iota 每遇到一次关键字const就会重置为0，所以多个const之间不会印象iota的取值
const (
	a1 = iota // 0
	a2 = iota // 1
	a3        // a3 声明也忽略了值，所以等于上一个值，也就是iota
	a4 = 100
	a5        // a5 忽略值，即取值会等于a4 100
	a6 = iota // 由于iota在a4 a5 声明过程中也在自增，所以a6=5
	_         //丢弃
)

const (
	c1, c2 = iota + 1, iota + 1 // 注意 iota自增的条件：const内每新增一行声明自增1，所以同一行内iota值不变
	c3, c4 = iota + 2, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
)

const (
	t  string = "rrr"
	t1 int    = 100
)

func main() {
	fmt.Println(n1, n2, n3) // 所有值都等n1 100
	fmt.Println(a1, a2, a3, a4, a5, a6)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println(t, t1)
}
