package main

// 类型断言
import "fmt"

type Usb interface {
	start()
	end()
}

type Phone struct {
	name string
}

func (p *Phone) start() {
	fmt.Println(p.name, " start")
}

func (p *Phone) end() {
	fmt.Println(p.name, " end")
}

func (p *Phone) call() {
	fmt.Println(p.name, " call")
}

type Camera struct {
	name string
}

func (c *Camera) start() {
	fmt.Println(c.name, " start")
}

func (c *Camera) end() {
	fmt.Println(c.name, " end")
}

type Computer struct {
}

func (c *Computer) Working(usb Usb) {
	usb.start()
	// 如果在Phone结构体的时候还需要调用Call，则需要使用类型断言
	if p, ok := usb.(*Phone); ok { // 注意如果类型是指针，这里断言也需要是指针
		p.call()
	}
	usb.end()
}

func main() {
	usbArr := [3]Usb{}
	usbArr[0] = &Phone{name: "Phone1"}
	usbArr[1] = &Phone{name: "Phone2"}
	usbArr[2] = &Camera{name: "Camera1"}
	fmt.Println(usbArr)

	var computer *Computer
	for _, v := range usbArr {
		computer.Working(v)
	}
}
