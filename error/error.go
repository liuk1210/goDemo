package main

import (
	"errors"
	"fmt"
)

func main() {
	exception()
	println("main继续执行。。。")
}

func exception() {
	//异常处理，defer recover
	defer func() {
		var err = recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//自定义错误，接收错误并终止程序
	err := errors.New("自定义错误")
	panic(err)
}
