package main

import (
	"fmt"
	"sort"
)

type Book struct {
	name   string
	sortNo int
}

type Books []Book

func (books Books) Len() int {
	return len(books)
}
func (books Books) Less(i, j int) bool {
	return books[i].sortNo > books[j].sortNo
}
func (books Books) Swap(i, j int) {
	tmp := books[i]
	books[i] = books[j]
	books[j] = tmp
}

func main() {
	var bs = Books{{name: "语文", sortNo: 2}, {name: "英语", sortNo: 3}, {
		name: "数学", sortNo: 1,
	}}
	fmt.Println(bs)
	sort.Sort(bs)
	fmt.Print(bs)
}
