package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "d:\\GoProjects\\src\\code.tjx.com\\study\\day01\\helloworld\\" // 转义
	fmt.Println(path)
	// 多行字符串
	s2 := `test
	test
	test` // ``会原样输出，如果不想添加多个转义符，可以使用``
	fmt.Println(s2)

	// 字符串长度
	fmt.Println(len(s2))

	// 字符串拼接
	name := "Im"
	word := "OK"
	ss := name + word
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s test %s", name, word) // Sprintf是格式化并返回
	fmt.Println(ss1)
	//  字符串分割
	ret := strings.Split(path, "\\")
	fmt.Printf("result:%v, Type:%T\n", ret, ret)
	// 拼接
	fmt.Println(strings.Join(ret, "+"))
	// 包含
	fmt.Println(strings.Contains(path, "Go"))
	// 前缀后缀
	fmt.Println(strings.HasPrefix(path, "d:"))
	fmt.Println(strings.HasSuffix(path, "test"))

	//
	s4 := "abcde"
	fmt.Println(strings.Index(s4, "c"))

	// 修改字符，不能直接修改
	testStr := "test@test"
	// testStr[4] = 'c'	// 错误
	// 需要先转成切片[]byte,或者[]rune再修改
	testByte := []byte(testStr)
	testByte[4] = 'c'
	testStr = string(testByte)
	fmt.Println(testStr)
	// 如果字符串含有英文和数字意以外的字符，其他字符可能占3个字节，譬如中文，不能转成[]byte，则需要使用rune
	testStr = "tjx谭健雄@test"
	testRune := []rune(testStr)
	testRune[2] = 'c'
	testRune[3] = '北'
	testStr = string(testRune)
	fmt.Println(testStr)

}
