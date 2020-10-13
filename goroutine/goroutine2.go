package main

import (
	"fmt"
	"time"
)

//找出1-n中所以的质数

//向管道numberChan中放入1-n的数据
func putNumber(numberChan chan int, n int) {
	for i := 1; i <= n; i++ {
		numberChan <- i
	}
	close(numberChan)
}

func calc(numberChan chan int, resultChan chan int, exitChan chan bool) {
	for {
		number, ok := <-numberChan
		if !ok { //如果numberChan中的数据读完了
			break
		}
		//是否质数
		flag := true
		for i := 2; i < number; i++ {
			if number%i == 0 {
				flag = false
				break
			}
		}
		if flag { //如果是则存放入结果管道中
			resultChan <- number
		}
	}
	exitChan <- true //向退出管道存入数据
}

//读取退出管道数据，阻塞等待所有计算协程计算完毕
func exit(exitChan chan bool, resultChan chan int, count int) {
	for i := 1; i <= count; i++ {
		<-exitChan
	}
	close(resultChan)
}

func main() {
	var now = time.Now().UnixNano()
	var numberChan = make(chan int, 100)
	var resultChan = make(chan int, 100)
	var exitChan = make(chan bool, 2)
	//启动协程 放入数据
	go putNumber(numberChan, 10000)
	var count = 12
	//启动12个协程去计算质数
	for i := 1; i <= count; i++ {
		//从numberChan中取出数据，计算是否质数，如果是则放入resultChan
		//如果numberChan中的数据读取完，则退出并想exitChan中写入true表示已完成
		go calc(numberChan, resultChan, exitChan)
	}
	//待所有协程完成后关闭结果管道
	go exit(exitChan, resultChan, count)
	for { //遍历结果数据
		v, ok := <-resultChan
		fmt.Print(v, " ")
		if !ok { //读取完毕
			break
		}
	}
	fmt.Println("\n计算完毕，耗时：", time.Now().UnixNano()-now)

}
