package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// map也是引用类型
func main() {
	var a map[string]int
	if a == nil {
		fmt.Println("a is nil") // a 还未初始化，是nil，不能赋值
	}
	a = make(map[string]int, 8)
	fmt.Println(a == nil) // 已经初始化，不是nil，可以赋值
	a["a"] = 100
	a["b"] = 200
	fmt.Printf("%#v\n", a)

	// 声明并赋值
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("%#v len:%d\n", b, len(b))

	// // 错误的添加键值对的方式
	// var c map[int]bool // 不初始化的情况下，是一个nil，相当于没有申请内存空间
	// c[100] = true      // 所以无法赋值，只有初始化之后才能赋值
	// fmt.Println(c)

	// 判断某个键值对是否存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["name1"] = 100
	scoreMap["name2"] = 200
	v, ok := scoreMap["name1"]
	fmt.Println(v, ok)
	v, ok = scoreMap["name3"]
	fmt.Println(v, ok)

	// 遍历
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}

	// 删除
	delete(scoreMap, "name1")
	fmt.Println(scoreMap)

	// 批量添加
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	// 有序遍历以上scoreMap
	// 1、先把key取出并进行排序
	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	// 2、按照排序后的key把值取出
	for _, k := range keys {
		fmt.Println(k, scoreMap[k])
	}

	// 元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8)
	// 注意如果每个元素都是map，则这样声明初始化只初始化了mapSlice，而里面的元素其实还没初始化，元素都是nil
	// [nil, nil, nil, nil, nil, nil ...]
	fmt.Println(mapSlice[0] == nil) // True
	// 因为里面的元素是nil，所以不能使用赋值mapSlice[0]["test"] = 100 这样的赋值
	mapSlice[0] = make(map[string]int, 1)
	mapSlice[0]["test"] = 100
	fmt.Println(mapSlice[0])

	// 值为切片的map
	var sliceMap = make(map[string][]int, 8) // 只完成了map的初始化，切片的初始化还未完成
	val, ok := sliceMap["zh"]                // 对于key不存在的错误可以这样处理
	if ok {
		fmt.Println(val)
	} else {
		sliceMap["zh"] = make([]int, 8, 8)
		sliceMap["zh"][0] = 100
		sliceMap["zh"][1] = 200
		sliceMap["zh"][2] = 300
	}
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}

	// 练习，统计单词出现的次数
	var sentences string = "how do you do"
	var words []string = strings.Split(sentences, " ")
	fmt.Println(words)
	var wordCount = make(map[string]int, len(words))
	for _, w := range words {
		c, ok := wordCount[w]
		if ok {
			wordCount[w]++
			fmt.Println(c)
		} else {
			wordCount[w] = 1
		}
	}
	for key, value := range wordCount {
		fmt.Println(key, value)
	}
}
