package main

import (
	"fmt"
	"time"
)

/*
在某些场景下我们需要同时从多个通道接收数据,通道在接收数据时,如果没有数据可以接收将会发生阻塞。为了应对这种场景，Go内置了select关键字，可以同时响应多个通道的操作。

select多路复用:
	类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。
	具体格式如下：
		select {
    		case <-chan1:
       			// 如果chan1成功读到数据，则进行该case处理语句
    		case chan2 <- 1:
       			// 如果成功向chan2写入数据，则进行该case处理语句
    		default:
       			// 如果上面都没有成功，则进入default处理流程
    	}
*/

//定时5秒
func s1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "s1"
 }

 //定时2秒
 func s2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "s2"
 }

 func test1(){
	 // 2个管道
	 output1 := make(chan string)
	 output2 := make(chan string)
	 // 跑2个子协程，写数据
	 go s1(output1)
	 go s2(output2)
	 // 用select监控
	 select {
		case s1 := <-output1:
			fmt.Println("s1=", s1)
		case s2 := <-output2:
			fmt.Println("s2=", s2)
	 }
 }

 func test2(){
	 // 创建2个管道
	 int_chan := make(chan int, 1)
	 string_chan := make(chan string, 1)
	 go func() {
		int_chan <- 1
	 }()
	 go func() {
		string_chan <- "hello"
	 }()
	 select {
		case value := <-int_chan:
			fmt.Println("int:", value)
		case value := <-string_chan:
			fmt.Println("string:", value)
	 }
 }

 func write(ch chan string) {
	for {
	   select {
			// 写数据
			case ch <- "hello":
				fmt.Println("write hello")
			default:
				fmt.Println("channel full")
	   }
	   time.Sleep(time.Millisecond * 500)
	}
 }

 func test3(){
	 // 创建管道
	 output1 := make(chan string, 10)
	 // 子协程写数据
	 go write(output1)
	 // 取数据
	 for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	 }
 }

 func main() {
	 //select基本用法
	 //select可以同时监听一个或多个channel，直到其中一个channel ready
	//  test1()

	 //如果多个channel同时ready，则随机选择一个执行
	//  for i:=0; i<888;i++{
	// 	test2()
	//  }

	 //可以用于判断管道是否存满
	 test3()
 }