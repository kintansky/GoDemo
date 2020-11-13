package main

import "fmt"

// 匿名字段
type Person struct {
	string // 匿名字段，匿名字段同类型的只能有一个
	int
}

func main() {
	p1 := Person{
		string: "tjx",
		int:    20,
	}
	fmt.Println(p1.string, p1.int)

}
