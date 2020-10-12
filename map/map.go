package main

import "fmt"

func main() {
	var a map[string]string = make(map[string]string)
	a["v1"] = "222"
	a["v2"] ="333"
	fmt.Println(a)

	//查找key
	val,hasKey := a["v1"]
	if hasKey {
		fmt.Println(val)
	}

	//遍历
	for key,val :=range a{
		fmt.Printf("key=%v,value=%v\n",key,val)
	}


	delete(a,"v1")
	fmt.Println(a)

	a = make(map[string]string)
	fmt.Println(a)
}
