package main

import "fmt"

func main() {
	// byte uint8的别名 ascii
	// rune int32的别名

	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T, c2:%T", c1, c2)

	s := "hello谭健雄"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i]) // 这样输出会出现错误，因为中文这些一个字符占多个字节，打印字符不能使用%v
	}
	// 正确做法
	for _, r := range s { // range 会判断每个字符占用的字节数
		fmt.Printf("%c\n", r)
	}
}
