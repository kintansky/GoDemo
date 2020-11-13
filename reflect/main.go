package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	// 不知道传进来的参数是什么类型
	// 1、可借助类型断言
	// 2、利用反射
	obj := reflect.TypeOf(x) // 返回的obj本身是一个反射类型指针*reflect.rtype
	fmt.Printf("%T, %v\n", obj, obj)
	fmt.Println(obj.Name(), obj.Kind()) // Name()对应的是类型名，Kind()对应的是种类
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%T, %v\n", v, v) // 返回的是一个reflect.Value类型，值是对应的值
	// 如果需要拿到对应类型的变量，通过以下方法
	k := v.Kind()
	switch k {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("%T, %v\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%T, %v\n", ret, ret)
	}
}

// 通过类型判断修改原来的值
func reflectSetValue(x interface{}) { // 修改必须传指针
	v := reflect.ValueOf(x)
	// Elem()
	k := v.Elem().Kind() // 因为v是指针，需要通过Elem()取到地址对应的值
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100) // 这里也需要使用Elem
	case reflect.Float32:
		v.Elem().SetFloat(1.2)
	}
}

type Cat struct {
}

type Dog struct {
}

func main() {
	var a float32 = 1.2
	reflectType(a)
	reflectValue(a)
	var b int8 = 10
	reflectType(b)
	reflectValue(b)
	// // 结构体
	// var c Cat
	// reflectType(c)
	// var d Dog
	// reflectType(d)
	// // type name
	var aaa int32 = 99
	reflectSetValue(&aaa)
	fmt.Println("After set: ", aaa) // aaa 的在值被修改
}
