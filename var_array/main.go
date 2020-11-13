package main

import "fmt"

func main() {
	var a [3]int
	var b [4]int
	fmt.Printf("%T, %v\n", a, a)
	fmt.Println(b)
	// 数组包含数据长度和类型，所以a和b是两种不同的类型不能互相赋值

	// 数组初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)
	// 自动推导长度初始化
	var boolArray = [...]bool{true, false, true}
	fmt.Println(boolArray)
	// 使用索引初始化
	var langArray = [...]string{1: "Golang", 3: "Python", 7: "Java"} // 可以不连续
	fmt.Println(langArray)

	// 数组的遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	for index, value := range cityArray { // for range 可以获得index和value两个值
		fmt.Println(index, value)
	}

	// 二位数组
	cityPairArray := [3][2]string{
		// 如果是多维数组只能再外层使用长度推导，内层不能使用长度推导[...][2]string{}
		{"北京", "西安"},
		{"上海", "杭州"},
		{"广州", "深圳"},
	}
	fmt.Println(cityPairArray)
	// 多维数组的遍历
	for index1, value1 := range cityPairArray {
		for index2, value2 := range value1 {
			fmt.Println(index1, index2, value2)
		}
	}
}
