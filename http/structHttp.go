package main

import (
	"fmt"
	"net/http"
)

type HandlerDemo struct {
}

func (handler *HandlerDemo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello1323235")
}

//结构体处理器
func main() {

	demo := HandlerDemo{}

	//方式一
	//http.Handle("/hello",&demo)
	//http.ListenAndServe(":8080",nil)

	//方式二
	//server := http.Server{
	//	Addr:":8080",
	//	Handler: &demo,
	//	ReadTimeout: 60*time.Second,
	//}
	//server.ListenAndServe()

	//方式三

	mux := http.NewServeMux()
	mux.Handle("/", &demo)
	http.ListenAndServe(":8080", mux)

}
