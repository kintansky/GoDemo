package main

// 接口是一种抽象的类型，无法实例化
// 只要结构体实现了接口下的 所有 方法，这个结构体就实现了这个接口，不需要指定
import "fmt"

type dog struct {
}

func (d dog) say() {
	fmt.Println("wang")
}

type cat struct {
}

func (c cat) say() {
	fmt.Println("miao")
}

type person struct {
	name string
} // person 不实现say()方法的话，无法传给da(),因为da()要求interface必须有一个say()

// 定义一个抽象类型，只要实现了say()方法的都可以称为sayer类型
type sayer interface {
	say()
	// move()	// 注意如果含有多个方法在里面，这些方法之间是且的关系，意味着传进来的类型都要含有这些方法才行，缺一不可
}

// 既要适配dog类型又要适配cat类型
func da(arg sayer) {
	arg.say()
}

type fish struct {
}

func (f *fish) say() { // 如果方法使用了指针接收者实现了接口，那么只有类型的指针才可以保存到类型的变量中
	fmt.Println("...")
}

func main() {
	// 可以直接传这个实例
	c1 := cat{}
	da(c1)
	// 也可以通过先声明接口在赋值
	var s sayer
	d1 := dog{}
	s = d1
	da(s)
	// p1 := person{
	// 	name: "tjx",
	// }
	// da(p1) // person没有say()方法，无法调用da()

	// 使用指针接收者
	f1 := &fish{}
	s = f1  // 接口可以接收指针
	s.say() // say()使用的是一个指针接收者，
	/*
		// 需要注意一点：
		f2 := fish{}	// 使用值类型
		s = f2 // 无法赋值，因为f2是fish类型的值不是指针，没有实现say()接口
		s.say()
	*/

}
