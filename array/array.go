package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//数组的几种定义方式
	var arr1 [2]int = [2]int{2, 3}
	var arr2 = [3]int{2, 3, 1}
	var arr3 = [...]float32{2, 3}
	var arr4 = [...]float32{1: 2, 0: 3}
	fmt.Println(arr4)

	//for range遍历数组
	for _, item := range arr1 {
		fmt.Printf("item的值：%v\n", item)
	}
	//for range 遍历数组
	for index, value := range arr2 {
		fmt.Printf("第%d个元素：%v\n", index, value)
	}
	fmt.Println("传统方式遍历数组")
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("第%d个元素：%v\n", i, arr3[i])
	}
	fmt.Printf("%p\n", &arr2)
	modify(arr2)
	fmt.Println(arr2) //没有修改
	realModify(&arr2)
	fmt.Println(arr2) //修改成功

	rand.Seed(time.Now().UnixNano())
	var arr5 [5]int
	for index,_ := range arr5{
		arr5[index] = rand.Intn(100)
	}
	fmt.Println(arr5)

	//二维数组
	var arr12 [2][3]int = [2][3]int{{1,2},{2,2,3}}
	fmt.Print(arr12)
	var arr13 [2][3]int = [...][3]int{{1,2},{2,2,3}}
	fmt.Print(arr13)

}
//值传递
func modify(arr [3]int) {
	arr[1] = 12
}
//引用传递
func realModify(arr *[3]int) {
	(*arr)[1] = 22
}
