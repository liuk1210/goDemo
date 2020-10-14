package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.133.1:8888")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conn.RemoteAddr().String())
	for {

		reader := bufio.NewReader(os.Stdin)

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		line = strings.Trim(line," \r\n")
		if line == "exit" {
			fmt.Println("退出")
			break
		}
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("发送了%v个字节\n", n)
	}

}
