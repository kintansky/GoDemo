package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

// 通过反射拿到结构体包含的方法
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("Have %d method", t.NumMethod()) // 获取方法的数量
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args) // 调用方法，如果方法含有参数，需要将参数转换成[]reflect.Value{}类型
	}
	// 当然也可以通过方法名来调用
	r, ok := t.MethodByName("Sleep") // 返回的是方法和Bool
	if ok {
		fmt.Println(r)
	}
	v.MethodByName("Sleep") // 返回的是值
}

func main() {
	stu1 := student{
		Name:  "tjx",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Printf("name %s, kind %v\n", t.Name(), t.Kind())

	// 通过反射获取结构体所有字段信息
	for i := 0; i < t.NumField(); i++ { // NumField返回字段的数量
		fieldObj := t.Field(i)
		fmt.Printf("name:%v, type:%v, tag:%v\n", fieldObj.Name, fieldObj.Type, fieldObj.Tag)
		fmt.Println("JSON TAG:", fieldObj.Tag.Get("json")) // 这样可以只取json的tag
	}

	// 根据名字取结构体的字段
	fieldObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v, type:%v, tag:%v\n", fieldObj2.Name, fieldObj2.Type, fieldObj2.Tag)
	}

	//
	printMethod(stu1)
}
