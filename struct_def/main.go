package main

import "fmt"

// 结构体定义
type person struct {
	name, city string
	age        int
}

// go语言没有构造函数，需要通过下面方式实现
func newPerson(name, city string, age int) *person {
	return &person{ // 通过返回指针，避免多重复制的开销，增加性能
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	// 实例化
	var p person
	p.name = "tjx"
	p.city = "zh"
	p.age = 20
	fmt.Printf("%#v\n", p)

	// 匿名结构体， 临时使用
	var user struct {
		name    string
		married bool
	}
	user.name = "tjx"
	user.married = true
	fmt.Printf("%#v\n", user)

	// 结构体的指针
	var p1 = new(person) // 通过new创建一个person指针
	fmt.Printf("%p\n", p1)
	(*p1).name = "testname"
	(*p1).city = "hh"
	(*p1).age = 30
	// Go底层实现了指针和类型的转换，所以也可以直接通过 p.name = "..."的方式赋值
	fmt.Printf("%#v\n", p1) // 是一个person指针：&main.person{name:"testname", city:"hh", age:30}

	// 通过结构体的地址进行实例化
	p3 := person{}          // 注意是花括号
	fmt.Printf("%T\n", p3)  // 这时候p3已经实例化是一个person类型main.person
	fmt.Printf("%#v\n", p3) // 实例化的时候内部变量都是对应的零值main.person{name:"", city:"", age:0}

	p4 := &person{}         // 通过地址实例化
	fmt.Printf("%T\n", p4)  // 这时候p3已经实例化是一个person指针*main.person
	fmt.Printf("%#v\n", p4) // 区别于上面，p4是一个person指针&main.person{name:"", city:"", age:0}

	// 结构体的初始化
	// 1、键值对初始化
	p5 := person{
		name: "tjx",
		age:  18,
		// 不一定所有字段都初始化，但是使用键值对初始化的时候各个字段后面都要协商逗号
	}
	fmt.Printf("%#v\n", p5)

	// 更常用的方法，因为结构体是值类型，使用指针性能会更好，避免复制
	p6 := &person{ // 初始化的时候给值p6赋值他的地址
		name: "tjx",
		age:  18,
		// 不一定所有字段都初始化，但是使用键值对初始化的时候各个字段后面都要协商逗号
	}
	fmt.Printf("%#v\n", p6)

	// 2、值列表初始化
	p7 := person{
		"tjx",
		"bj",
		18,
		// 使用值列表的方式初始化，必须全匹配
	}
	fmt.Printf("%#v\n", p7)

	// 构造函数
	p8 := newPerson("tjx", "bj", 20)
	fmt.Printf("%#v\n", p8)

}
