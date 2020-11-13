package main

import (
	import "github.com/atotto/clipboard"
)

func main(){
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}