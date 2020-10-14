package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@tcp(192.168.133.222:3306)/go")
	if err != nil {
		panic(err)
	}
	row := db.QueryRow("show databases")
	if err!=nil{
		fmt.Println(err)
	}
	var col string
	row.Scan(&col)
	fmt.Print(col)

}
