package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("开始监听8888端口。。。")
	//tcp协议
	//监听8888端口
	listener,err := net.Listen("tcp","0.0.0.0:8888")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer listener.Close()
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println(err)
		}
		go connect(conn)
		fmt.Println("remoteAddr：",conn.RemoteAddr().String())
	}

}

func connect(conn net.Conn){
	defer conn.Close()
	for{
		buf := make([]byte,1024)
		fmt.Println("正在等待客户端发送消息,地址为：",conn.RemoteAddr().String())
		//等待客户端通过conn发送信息
		//如果客户端没有write，则阻塞
		n,err := conn.Read(buf)
		if err !=nil{
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
