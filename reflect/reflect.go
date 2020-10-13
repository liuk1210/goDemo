package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	sub := c.Num1 - c.Num2
	fmt.Printf("%v完成了减法运算，%v-%v=%v", name, c.Num1, c.Num2, sub)
}

//反射机制
func Reflect(cal interface{}) {
	typ := reflect.TypeOf(cal)
	val := reflect.ValueOf(cal)

	//遍历字段
	num := val.Elem().NumField()
	for i := 0; i < num; i++ {
		fmt.Println(typ.Elem().Field(i).Name)
	}

	//设置参数值
	val.Elem().Field(0).SetInt(8)
	val.Elem().Field(1).SetInt(3)

	//参数设置，反射调用方法
	var param []reflect.Value
	param = append(param, reflect.ValueOf("tom"))
	val.Method(0).Call(param)
}

func main() {
	cal := Cal{}
	Reflect(&cal)
}
