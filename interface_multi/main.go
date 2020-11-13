package main

import "fmt"

type AInterface interface {
	yell()
}

type BInterface interface {
	move()
}

// CInterface继承AInterface和BInterface接口
type CInterface interface {
	AInterface
	BInterface
}

// Cat 实现了ABC所有接口
type Cat struct {
}

func (c *Cat) yell() {
	fmt.Println("miaomiaomiao")
}

func (c *Cat) move() {
	fmt.Println("move")
}

// DOG 只能实现AInterface，B和CInterface都不能实现
type Dog struct {
}

func (d Dog) yell() {
	fmt.Println("wangwangwang")
}

func main() {
	c := Cat{}
	// 注意1
	// 赋值给接口的时候，需要关注对应的方法实现是使用了指针还是值传递
	// 譬如c的所有方法我都使用指针传递的话
	// 接下来赋值给接口的时候，也要用Cat指针赋值给接口
	var ai AInterface = &c
	var bi BInterface = &c
	var ci CInterface = &c
	ai.yell()
	bi.move()
	ci.yell()
	ci.move()

	d := Dog{}
	// 注意2
	// DOG的方法使用了值传递，所以赋值给接口的时候应该使用值
	ai = d
	// bi = &d	// 这里就会提示错误，因为DOG没有实现BInterface
	// ci = &d	// 这里就会提示错误，因为DOG没有实现CInterface
	ai.yell()
	// 注意3
	// 所以结构体里面的方法应该统一接收者类型，否则接口需要定义多个

}
