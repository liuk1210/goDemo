package main

import "fmt"

//结构体继承
type A struct{
	name string
}
//通过内嵌的方式实现继承，继承时可访问所有父结构体的方法与属性
type B struct {
	A
	code string	//小写非本包无法访问
}
func (b B) SetCode(code string){
	b.code = code
}

func main() {
	var b B
	b.name="bb"
	b.code="b1"
	fmt.Print(b)
}
