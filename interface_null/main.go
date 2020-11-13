package main

import "fmt"

/*
空接口：内部没有定义方法，就是一个空接口
空接口因为没有方法，所以可以存任意值
type xxx interface {
}
所以空接口一般不全局定义，只要在使用的时候定义即可
*/

// 空接口的应用
// 1、空接口作为函数的参数
// 2、空接口作为map的values

func main() {
	var x interface{} // 定义一个空接口
	x = "hello"
	fmt.Println(x)
	x = 100
	fmt.Println(x)
	x = false
	fmt.Println(x)

	// 应用
	var m = make(map[string]interface{}, 16) // 使用空接口扩展map的功能，使得values可以接收任何类型的值
	m["name"] = "tjx"
	m["age"] = 20
	m["hobby"] = []string{"篮球", "足球"}
	fmt.Printf("%#v\n", m)

	// 类型断言
	ret, ok := x.(string) // 猜x是一个string
	if !ok {
		fmt.Println("guess again")
	} else {
		fmt.Println(ret)
	}

	// 使用switch做类型断言
	switch t := x.(type) {
	case string:
		fmt.Println("type string", t)
	case bool:
		fmt.Println("type bool", t)
	case int:
		fmt.Println("type int", t)
	default:
		fmt.Println("unknow", t)
	}

	// 使用null interface 也可以创建不受类型限制的array
	var l = make([]interface{}, 0, 10)
	l = append(l, 0)
	l = append(l, "test")
	fmt.Println(l)

}
