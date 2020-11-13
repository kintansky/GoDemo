package main

import "fmt"

func main() {
	// swith 只要符合就会退出，不会再像C一样需要break，也就可以强制使用fallthrough来适应C的语法
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
		fmt.Println("haha")
	case 4:
		fmt.Println("无名指")
	default: // 其他情况
		fmt.Print("无效")
	}
	// case内跟多个值
	num := 5
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}
	//case判断语句
	num = 5
	switch { // 注意这里如果case使用的是表达式，switch后不能加其他内容
	case num%2 != 0:
		fmt.Println("奇数")
	case num%2 == 0:
		fmt.Println("偶数")
	case num == 100:
		// case 100:	// 错误写法，因为本switch已经使用了表达式，所以这里也只能跟表达式
		fmt.Println("大")
	}
}
