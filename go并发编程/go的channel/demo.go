package main

import (
	"fmt"
)

/*
单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。使用共享内存进行数据交换会在不同的goroutine中容易发生竞态问题。因此要使用互斥量进行加锁以保证数据交换的正确性，但会造成性能问题。

Go语言的并发模型是CSP，提倡通过通信共享内存，而不是通过共享内存来实现通信。

Go语言中的通道(channel):
	channel是一种特殊的类型(也是引用类型), 声明channel的时候需要为其指定元素类型
	channel像一个传送带或者队列，遵循先入先出的规则，保证收发数据的顺序
	channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制

创建channel的格式：
	1、先声明、后初始化（分配内存）
		例子：var ch chan int //声明一个int类型的channel
			 ch = make(chan int)  //初始化已声明的channel
    2、通过make直接创建
		例子：ch := make(chan int) //创建int类型的channel

channel操作:
	通道有发送、接收和关闭三种操作。发送和接收都使用<-符号
	1、向channel发送数据: ch <- 10
	2、从channel接收数据: i := <- ch
	3、关闭channel: close(ch) //调用内置的close函数来关闭通道
		关于关闭通道需要注意的事情:
			通道是可以被垃圾回收机制回收的,
			只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道,
			它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
		关闭后的通道有以下特点：
			1.对一个关闭的通道再发送值就会导致panic。
			2.对一个关闭的通道进行接收会一直获取值直到通道为空。
			3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
			4.关闭一个已经关闭的通道会导致panic。

注意: 如果创建的是无缓冲区的channel，那么在同一个goroutine任务重发送、接受数据时会阻塞，报死锁异常
	  无缓冲区channel: ch := make(chan int)
	  有缓冲区channel: ch := make(chan int,10) // 10就是缓冲区的大小

*/

func recive(ch chan int){
	i := <- ch
	fmt.Println("无缓冲区channel接收值=",i)

}

func reciveAndClose(){
	ch1 := make(chan int)
    ch2 := make(chan int)
    // 开启goroutine将0~100的数发送到ch1中
    go func() {
        for i := 0; i < 100; i++ {
            ch1 <- i
        }
        close(ch1)
    }()
    // 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
    go func() {
        for {
            i, ok := <-ch1 //方式一判断channel关闭 通道关闭后再取值ok=false
            if !ok {  
                break
            }
            ch2 <- i * i
        }
        close(ch2)
    }()
    // 在主goroutine中从ch2中接收值打印
    for i := range ch2 { //方式二判断channel关闭： 通道关闭后会退出for range循环
        fmt.Println(i)
    }
}

func main() {
	//声明channel
	var ch chan int  //声明一个int类型的channel
	fmt.Println(ch) //声明是没有分配内存的，所以是nil
	//声明后的channel需要使用make函数初始化之后才能使用，或者直接通过make函数创建channel
	ch = make(chan int) //make初始化已声明的channel
	fmt.Println(ch)
	ch1 := make(chan string,5) //创建string类型的channel,带缓冲区的channel
	fmt.Println(ch1)

	//channel操作,通道有发送、接收和关闭三种操作。发送和接收都使用<-符号
	//注意：发送和接受的channel是在main函数(主goroutine)，因此使用无缓冲区channel会报死锁异常
	// 解决方式一: 使用带缓冲区的channel
	//发送数据到channel
	ch1 <- "hello channel"
	//从channel接收数据
	s := <- ch1
	fmt.Println(s)
	//解决方式而: 启用一个goroutine接收值,必须先启动接收，然后才发送数据
	go recive(ch) 
	ch <- 100
	//关闭channel,调用内置的close函数来关闭通道
	close(ch)
	close(ch1)

	//循环从channel取值，以及判断channel是否关闭
	reciveAndClose()
	
	/*
	单向channel:
		1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
		2.<-chan int是一个只能接收的通道，可以接收但是不能发送。   
	*/
	oneWayChannel()
}

func oneWayChannel(){
	ch1 := make(chan int)
    ch2 := make(chan int)
    go counter(ch1)
    go squarer(ch2, ch1)
    printer(ch2)
}

//只能向channel发送数据
func counter(out chan<- int) {
    for i := 0; i < 100; i++ {
        out <- i
    }
    close(out)
}
//参数一：只能向channel发送数据
//参数二：只能从channel接收数据
func squarer(out chan<- int, in <-chan int) {
    for i := range in {
        out <- i * i
    }
    close(out)
}
//只能从channel接收数据
func printer(in <-chan int) {
    for i := range in {
        fmt.Println(i)
    }
}
