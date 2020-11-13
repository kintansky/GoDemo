package main

import "fmt"

// TypeJudge 针对不同的类型就行switch判断断言
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("index %d: %T, %v\n", index, x, x)
		case float32:
			fmt.Printf("index %d: %T, %v\n", index, x, x)
		case float64:
			fmt.Printf("index %d: %T, %v\n", index, x, x)
		case int, int32, int64:
			fmt.Printf("index %d: %T, %v\n", index, x, x)
		case string:
			fmt.Printf("index %d: %T, %v\n", index, x, x)
		default:
			fmt.Printf("unknown\n")
		}
	}
}

func main() {
	var arr = make([]interface{}, 5)
	arr[0] = 1
	arr[1] = 1.2
	arr[2] = "test"
	arr[3] = false
	TypeJudge(arr...)

	TypeJudge(1, 2.0, "ttt")
}
