package main

// 函数、匿名函数
import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func sayHello2(name string) {
	fmt.Println("hello" + name)
}

func intSum(a int, b int) int {
	ret := a + b
	return ret
}

func intSum2(a int, b int) (ret int) {
	// 注意如果返回值指明了返回值名称，则函数体内可以不能再声明这个返回值
	ret = a + b // 这里没有再声明
	return ret  // 返回值也可以直接简写return，这样会函数会找到ret返回，但不推荐简写
}

// 接收可变参数
// 另外，GO语言没有默认参数
func intSum3(a ...int) int { // 如果传参含有固定参数也含有可变参数，可变参数要放最后
	fmt.Printf("%v, %T\n", a, a) // 打印类型发现其实通过这样传进来的a是一个切片
	ret := 0
	for _, v := range a {
		// fmt.Println(v)
		ret += v
	}
	return ret
}

// 多返回值
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return sum, sub
}

func main() {
	sayHello2("tjx")
	s := intSum3(10, 20, 30)
	fmt.Println(s)

	x, y := calc(100, 200)
	fmt.Println(x, y)

	// 匿名函数，区别：没有函数名
	niming := func(x, y int) {
		fmt.Println(x + y)
	}
	// 匿名函数的执行
	niming(10, 20)
	// 还可以直接定义的时候执行
	func() {
		fmt.Println("niming")
	}()
}
