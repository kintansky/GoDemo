package main

import "fmt"

// slice, map, chan, interface 引用类型默认都是引用传递
func test(m map[string]int) {
	fmt.Printf("%p\n", m)
	m["c"] = 3
}

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Printf("%p\n", m)
	test(m)
	fmt.Printf("%v", m)
}
