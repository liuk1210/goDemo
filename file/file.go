package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	CopyFile("d:/1.txt","d:/2.txt")
}

func CopyFile(path1,path2 string) (written int64, err error) {
	file1,err := os.Open(path1)
	if err !=nil{
		fmt.Println(err)
		panic(err)
	}
	defer file1.Close()
	reader := bufio.NewReader(file1)

	file2,err :=os.OpenFile(path2,os.O_CREATE|os.O_WRONLY,0666)
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
	defer file2.Close()
	writer := bufio.NewWriter(file2)

	return io.Copy(writer,reader)

}

//生成并写入文件
func writeFile(){
	file,err :=os.OpenFile("d:/1.txt",os.O_CREATE|os.O_WRONLY,0666)
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("测试写入文件")
	writer.Flush()
}

//判断文件是否存在
func fileExists(path string) (bool,error){
	_,err := os.Stat(path)
	if err==nil{
		return true,nil
	}else if os.IsNotExist(err){
		return false,nil
	}
	return false, err
}

//一次性读取文件
func readFile(){
	file,err := ioutil.ReadFile("D:/GitHub/demo/map/map.go")
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(file))
}

//缓冲区读取大文件
func buffReader()  {
	file,err := os.Open("D:/GitHub/demo/map/map.go")
	if err !=nil{
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		line,e := reader.ReadString('\n')
		if e ==io.EOF{
			break
		}
		fmt.Println(line)
	}
	fmt.Println("读取完毕。")
}
