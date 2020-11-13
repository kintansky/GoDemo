package main

import "fmt"

// 切片是引用类型
func main() {
	// 切片的定义,切片实际还是一个数组
	var a []string // 但是注意这种方法，如哦个不适用make进行初始化，会提示nil，未分配内存，导致无法进行下一步操作
	var b []int
	var c = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 或者使用数组初始化
	a1 := [5]int{55, 56, 57, 58, 59}

	d := a1[1:4]
	fmt.Printf("%T, %v\n", d, d)

	// make创建
	d1 := make([]int, 5, 10) // 长度为5，最大容量为10，如果不指定容量，默认等于长度，注意使用make的时候其实已经初始化了
	fmt.Println(d1)
	fmt.Println(len(d1))
	fmt.Println(cap(d1)) // 获取容量

	// 切片之间不能使用== 来判断所有元素是否相等，切片只能和nil空值比较
	a1 = [5]int{55, 56, 57, 58, 59} // 但是数组可以直接比较
	a2 := [5]int{54, 56, 57, 58, 59}
	if a1 == a2 {
		fmt.Println("a1==a2")
	}
	// 如果要判断一个切片是否为空，是通过长度判断的，长度为0，单仍不等于nil
	var d2 []int // 只有只声明未初始化的时候即未申请内存才会是nil，包括var d2 []int{} 也不是nil
	if d2 == nil {
		fmt.Println("d2 is nil")
	}
	var d3 = make([]int, 0) // 仍然不等于nil
	if d3 == nil {
		fmt.Println("d3 is nil")
	} else {
		fmt.Println("d3 not nil")
	}
	// 遍历和数组一样
	// 切片的扩容
	var d4 []int
	d4 = append(d4, 10)
	fmt.Println(d4)
	for i := 0; i < 10; i++ {
		d4 = append(d4, i+10)
		fmt.Printf("%v, cap:%d, len:%d\n", d4, cap(d4), len(d4))
	}
	// 批量追加
	d4 = append(d4, 1, 2, 3, 4)
	fmt.Println(d4)
	// 追加切片,但无法追加数组，否则弹出类型出错
	d4Add := []int{100, 100, 100}
	d4 = append(d4, d4Add...) // ...表示将切片中的元素添加进去
	fmt.Println(d4)
	//如果要追加数组，可以先把数组转成切片
	d5Add := [...]int{999, 999}
	d4 = append(d4, d5Add[:]...) // d5Add[:]从头且到尾等价于d5Add[0:len(d5Add)]
	fmt.Println(d4)

	// 切片的copy
	d5 := []int{1, 2, 3, 4, 5, 6}
	e := make([]int, 6)
	e1 := e     // 如果使用这种方式赋值切片，他们底层实际上使用的是同一个数组
	copy(e, d5) // copy则是通过复制拷贝的方式，底层不是同一个数组
	e[0] = 100
	fmt.Println(d5) // 修改e不影响d5
	fmt.Println(e)
	fmt.Println(e1) // e1 仍然会等于修改后的e

	// 切片的删除
	d6 := []string{"bj", "sh", "sz", "hz"}
	// 切片没有直接的删除操作，但可以通过append来实现
	d6 = append(d6[0:1], d6[2:]...) // 删除index=1的元素 s = append(s[0:index], s[index+1:]...)
	fmt.Println(d6)
	var s = make([]string, 5, 10) //[     ],其实是5个" "
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		s = append(s, fmt.Sprintf("%v", i)) // 因为使用的是append所以会在尾插进去
	}
	fmt.Println(s)
	// 切片的修改要使用index，不能使用value
	// for index, _ := range s {
	// 	s[index] = "mod"
	// }
}
