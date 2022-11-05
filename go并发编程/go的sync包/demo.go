package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
sync.WaitGroup: Go语言中可以使用sync.WaitGroup来实现并发任务的同步,类似于java的栅栏
	sync.WaitGroup有以下几个方法：
		方法名								 功能
		(wg * WaitGroup) Add(delta int)		计数器+delta
		(wg *WaitGroup) Done()				计数器-1
		(wg *WaitGroup) Wait()				阻塞直到计数器变为0

		sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。
		例如:
			当我们启动了N 个并发任务时，就将计数器值增加N。
			每个任务完成时通过调用Done()方法将计数器减1。
			通过调用Wait()来等待并发任务执行完。
			当计数器值为0时，表示所有并发任务已经完成。

	注意: sync.WaitGroup是一个结构体，传递的时候要传递指针。


sync.Once: 在高并发的场景下只执行一次，例如只加载一次配置文件,单例模式(懒汉模式),
		   sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成
	sync.Once只有一个Do方法:
		func (o *Once) Do(f func()) {}
		//注意：如果要执行的函数f需要传递参数就需要搭配闭包来使用


sync.Map:
	Go语言中内置的map(var m = make(map[string]int))不是并发安全的,
	sync.Map是Go语言提供的并发安全版map，不用像内置的map一样使用make函数初始化就能直接使用。
	同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法
*/

//初始化配置函数
var icosn map[string]string
func loadIcons() {
    icosn = map[string]string{
        "left":  "left",
        "up":    "up",
        "right": "right",
        "down":  "down",
    }
}

/*
调用此方法获取配置值，如果配置没有初始化则先初始化
注意：当多个goroutine并发调用时不是并发安全的
	1、多个goroutine同时判断配置为nil，则可能会初始化多遍
	2、go编译器会做指令重排，将初始化配置函数重排成如下顺序：
		func loadIcons() {
			icons = make(map[string]image.Image)
			icons["left"] = loadIcon("left.png")
			icons["up"] = loadIcon("up.png")
			icons["right"] = loadIcon("right.png")
			icons["down"] = loadIcon("down.png")
		} 
		当一个goroutine执行了icons = make(map[string]image.Image)时，
			1、配置变量已经初始化完成，但后续的配置赋值并没有进行
			2、这时后面的goroutine判断icons不为nil，然后做取值操作，这时值是nil，
			   也就是配置变量分配了内存，但数据没准备好就开始操作

*/

func onceDemoOne(key string) string{
	if icosn == nil {
		loadIcons()
	}
	return icosn[key]
}

/*
使用sync.Once的Do方法解决上述问题
*/
var loadIconsOnce sync.Once
func onceDemoTwo(key string) string{
	loadIconsOnce.Do(loadIcons) //Do方法传入初始化配置函数
	return icosn[key]
}

/*
sync.Map
*/
var m = sync.Map{}
func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func(n int) {
            key := strconv.Itoa(n)
            m.Store(key, n)
            value, _ := m.Load(key)
            fmt.Printf("k=:%v,v:=%v\n", key, value)
            wg.Done()
        }(i)
    }
    wg.Wait()
} 