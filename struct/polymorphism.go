package main

import "fmt"

type Usb interface {
	run()
	stop()
}

//结构体只要实现了接口中定义的所有方法即视为实现了该接口
type Mouse struct {
}

func (m Mouse) run() {
	fmt.Println("mouse running ... ")
}
func (m Mouse) stop() {
	fmt.Println("mouse is stop")
}
func (m Mouse) click() {
	fmt.Println("mouse click")
}

type Keyboard struct {
}

func (m Keyboard) run() {
	fmt.Println("Keyboard running ... ")
}
func (m Keyboard) stop() {
	fmt.Println("Keyboard is stop")
}

type Computer struct {
}

//多态参数
func (c Computer) working(usb Usb) {
	usb.run()
	usb.stop()
}

func main() {
	c := Computer{}
	var usb [2]Usb //多态数组
	usb[0] = Mouse{}
	usb[1] = Keyboard{}
	for _, u := range usb {
		//类型断言
		if t, ok := u.(Mouse); ok {
			t.click()
		}
		c.working(u)
	}
}
