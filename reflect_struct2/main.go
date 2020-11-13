package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func (p *Person) show() {
	fmt.Println(p.name, p.age)
}

func main() {
	var a interface{}
	a = &Person{"tjx", 20}
	v := reflect.ValueOf(a) // 结果Person{"tjx", 20}类型reflect.Value
	fmt.Printf("%T, %v\n", v, v)
	// v.Type() 相当于reflect.TypeOf(a), 结果main.Person类型*reflect.rtype
	// v.Kind() = 系统类型就是struct这些

	t := reflect.TypeOf(a) // 结果main.Person类型*reflect.rtype
	fmt.Printf("%T, %v\n", t, t)
	trueType := t.Elem()
	fmt.Printf("%T, %v\n", trueType, trueType)
	ptrValue := reflect.New(trueType)
	fmt.Printf("%T, %v\n", ptrValue, ptrValue)
	// 下面两者效果一样
	trueValue := ptrValue.Elem()
	fmt.Printf("%T, %v\n", trueValue, trueValue)
	indirectValue := reflect.Indirect(ptrValue)
	fmt.Printf("%T, %v\n", indirectValue, indirectValue)

}
