package main

import (
	"fmt"
	"net/http"
)
//函数处理器
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"",r.URL.Path)
}
func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}
