package main

import "fmt"

type Address struct {
	Province   string
	City       string
	UpdateTime string // 同名字段
}

type Email struct {
	Addr       string
	UpdateTime string // 同名字段
}

type Person struct {
	Name   string
	Gender string
	Age    int
	// Address Address // 结构体嵌套，访问Address的字段需要通过p1.Address.Province, 理解为Person包含Address
	Address // 也可以作为匿名字段嵌套进来，使用匿名字段嵌套后可以直接访问p1.Province访问，不需要p1.Address.Province，理解为Person继承Address
	Email
}

func main() {
	p1 := Person{
		Name:   "tjx",
		Gender: "M",
		Age:    20,
		Address: Address{
			Province:   "GD",
			City:       "FS",
			UpdateTime: "2019-01-01",
		},
		Email: Email{
			Addr:       "CC",
			UpdateTime: "2020-12-12",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Address.Province) // 正常嵌套进来的只能通过这种方式访问
	fmt.Println(p1.Province)         // 通过匿名结构体嵌套进来可以通过这种方式访问

	// 注意如果多个嵌套的结构体含有同名字段，只能通过详细的嵌套访问方式
	fmt.Println(p1.Address.UpdateTime)
	fmt.Println(p1.Email.UpdateTime)

}
