package main

import (
	"fmt"
	"sync"
	// "time"
)

/*

在java/c++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个的任务，同时需要自己去调度线程执行任务并维护上下文切换。

那么能不能有一种机制，程序员只需要定义很多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行呢？
	Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程。
	但goroutine是由Go的运行时调度和管理的，Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。
	Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。

使用goroutine: 在Go语言编程中不需要去自己写进程、线程、协程
	当需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，
	在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。
	一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。

goroutine与线程:
	OS线程一般都有固定的栈内存(通常为2MB),
	goroutine栈在刚创建时只有很小的栈(典型情况下2KB)，所以在Go语言中一次创建十万左右的goroutine也是可以的。
	goroutine的栈不是固定的，可以按需增大和缩小，goroutine的栈大小限制可以达到1GB。

goroutine调度:
	GPM是Go语言运行时（runtime）层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。
	1.G: 就是goroutine，里面除了存放本goroutine信息外 还有与所在P的绑定等信息。
	2.P: 管理着一组goroutine队列
		 P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），
		 P会对自己管理的goroutine队列做一些调度,
		 	比如:暂停占用CPU时间较长的goroutine
			 	运行后续的goroutine等等
		 		当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。
	3.M: 是Go运行时对操作系统内核线程的虚拟
		 M与内核线程一般是一一映射的关系， 一个groutine最终是要放到M上执行的；

	P与M一般也是一一对应的。他们关系是： P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G 挂载在新建的M上。当旧的G阻塞完成或者认为其已经死掉时 回收旧的M。

	P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。 在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。

	单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。 其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能。


*/

func hello(){
	fmt.Println("Hello Goroutine")
}

//sync.WaitGroup类似于java中的栅栏，用来同步线程任务（这里是goroutine协程）
var wg sync.WaitGroup
func say(i int){
	defer wg.Done() // goroutine结束就登记-1
    fmt.Println("say Goroutine!", i)
}

//Go程序就会为main()函数创建一个默认的goroutine。
func main() {
	//启动一个goroutine执行hello方法
	go hello()
	fmt.Println("main goroutine")
	//当main()函数(默认的goroutine)返回的时候，所有在main()函数中启动的goroutine会一同结束
	// time.Sleep(time.Second)	//演示goroutine，使用睡眠（原理与java类似）


	//多次执行发现每次打印的数字的顺序都不一致。因为10个goroutine是并发执行的，而goroutine的调度是随机的
	for i:=0; i<10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go say(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}