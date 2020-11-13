package main

// 结构体的方法
import "fmt"

// Person is
type Person struct {
	name string
	age  int8
}

func newPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 作为Person的一个方法
func (p Person) Dream() {
	// 方法属于具体类型，这里Dream属于Person
	// 与函数的区别，函数不属于任何类型
	fmt.Printf("%s's Dream is\n", p.name)
}

// SetAge 使用是一个指针作为接收者，因为函数都是拷贝赋值，所以只能使用指针这种修改的操作才能成功
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

/* 什么情况下使用指针接收者
1、需要修改接收者的值
2、接收者是拷贝代价比较大的对象
3、保证一致性，如果其中一个方法使用了指针接收者，其他方法尽量也使用指针接收者
*/

func main() {
	p1 := newPerson("tjx", int8(20))
	p1.Dream() // 或者(*p1).Dream()

	fmt.Println(p1.age)
	p1.SetAge(int8(30))
	fmt.Println(p1.age)

}
