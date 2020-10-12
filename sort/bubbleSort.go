package main

import "fmt"

func main() {
	var array = []float64{1.2, 1.1, 3.2, 5.9, 0.3}
	fmt.Println("排序前：", array)
	BubbleSort(array)
	fmt.Println("排序后：", array)
}

//冒泡排序法
func BubbleSort(array []float64) {
	//进行array长度-1轮循环
	for i := 0; i < len(array)-1; i++ {
		//每轮循环找出最大的值放到最后
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] { //大于后面一个元素时 交换位置
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}
	}
}
