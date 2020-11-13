package main

// 结构体继承
import "fmt"

// 利用结构体的嵌套实现继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s move\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // Animal 做匿名字段嵌套进来,这样Dog就可以继承了Animal的方法
}

func (d *Dog) wang() {
	fmt.Printf("%s wang\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "dog01",
		},
	}
	d1.wang()
	d1.move()
}
