package main

import (
	"flag"
	"fmt"
	"os"
)
//读取命令行参数
func main() {
	fmt.Println(os.Args)
	var user string
	flag.StringVar(&user,"user","user","用户名为空")
	flag.Parse()
	fmt.Println(user)
}
