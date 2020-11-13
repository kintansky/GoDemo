package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	ID   int
	Name string
}

func newStudent(id int, name string) student { // 注意如果这里返回的使用指针，下面的也要更改为指针
	return student{ // 注意如果这里返回的使用指针，下面的也要更改为指针
		ID:   id,
		Name: name,
	}
}

type class struct {
	// 如果结构体中的字段名首字母小写，那么只能被当前包访问，其他包包括导入进来的json包，也不能访问
	// 如果部分项目需要指定字段名，可以通过tag实现，利用tag的时候，传给对应包的时候名字会做对应改变
	Title    string    `json:"title" db:"t"`
	Students []student // 注意如果这里返回的使用指针，下面的也要更改为指针
}

func main() {
	c1 := class{
		Title:    "class01",
		Students: make([]student, 0, 20), // 注意如果这里返回的使用指针，下面的也要更改为指针
	}

	for i := 0; i < 10; i++ {
		c1.Students = append(c1.Students, newStudent(i, fmt.Sprintf("student%02d", i)))
	}
	fmt.Printf("%#v\n", c1)

	// json序列化
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json fail, err:", err)
		return
	}
	fmt.Printf("%T\n", data) //
	fmt.Printf("%s\n", data) //

	// json 反序列化
	jsonStr := `
	{"Title":"class01","Students":[{"ID":0,"Name":"student00"},{"ID":1,"Name":"student01"},{"ID":2,"Name":"student02"}]}
	`
	// 但是要注意：如果jsonStr的id是string类型，和struct无法对应的，也不可以正常反序列化
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("unmarshal fail", err)
		return
	}
	fmt.Printf("%#v", c2)
}
