package main

func main() {
	var intChan = make(chan int, 2)
	var exitChan = make(chan bool)

	//读取和写入在终端输出时可能会出现 先输出 read后输入write的现象
	//其实自始至终intChan中都是先写后读的
	//产生该现象的原因主要是两个协程是并发/并行执行的
	//由于write协程在写入intChan之后再执行输出内容的代码
	//假设read协程在write协程写入完成的瞬间就已经获取到了write协程写入的内容，
	//而且之后抢先获得cpu的执行权力，先一步输出了read，就会产生先读后写的假象
	go write(intChan)
	go read(intChan, exitChan)

	//阻塞直到exitChan关闭
	<-exitChan
}

func write(intChan chan int) {
	for i := 1; i <= 10; i++ {
		intChan <- i
		println("write:", i)
	}
	close(intChan)
}

func read(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		println("read:", v, "_", ok)
	}
	close(exitChan)
}
