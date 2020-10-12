package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string
}
//反序列化
func main() {
	str := "{\"Name\":\"测试\"}"
	var a A
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.Name)
}
