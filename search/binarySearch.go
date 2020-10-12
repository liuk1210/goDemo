package main

import "fmt"

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}
	binarySearch(&array, 0, len(array)-1, 2)
}

//二分查找，前提是array已经是有序的数组
func binarySearch(array *[7]int, left int, right int, searchValue int) {
	if left > right {
		fmt.Println("没有找到", searchValue)
		return
	}
	middle := (left + right) / 2
	fmt.Printf("left=%v,right=%v,middle=%v\n", left, right, middle)
	if searchValue > (*array)[middle] {
		binarySearch(array, middle+1, right, searchValue)
	} else if searchValue < (*array)[middle] {
		binarySearch(array, left, middle-1, searchValue)
	} else {
		fmt.Printf("已找到%v，下标为%v\n", searchValue, middle)
	}
}
