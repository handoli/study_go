package main

import (
	"fmt"
	"sync"
	"time"
)

/*
有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）

互斥锁:
	是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。其他的goroutine则在等待锁；当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。Go语言中使用sync包的Mutex类型来实现互斥锁。

读写互斥锁:
	互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。

	读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
*/

//这里需要传指针保证是同一个变量、锁、栅栏对象
func add(x *int64,lock *sync.Mutex,wg *sync.WaitGroup) {
    for i := 0; i < 5000; i++ {
        lock.Lock() // 加锁
        *x = *x + 1
        lock.Unlock() // 解锁
    }
    wg.Done()
}
func lockDemo(){
	var wg sync.WaitGroup
	var lock sync.Mutex
	var x int64	
	wg.Add(2)
    go add(&x,&lock,&wg)
    go add(&x,&lock,&wg)
	wg.Wait()
	fmt.Println(x)
}


func write(x *int64,rwlock *sync.RWMutex,wg *sync.WaitGroup) {
    rwlock.Lock() // 加写锁
    *x = *x + 1
    time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
    rwlock.Unlock()                   // 解写锁
    wg.Done()
}

func read(x *int64,rwlock *sync.RWMutex,wg *sync.WaitGroup) {
    rwlock.RLock()               // 加读锁
    time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
    rwlock.RUnlock()             // 解读锁
    wg.Done()
}
func rwLockDemo(){
	start := time.Now()
	var (
		x      int64
		wg     sync.WaitGroup
		rwlock sync.RWMutex
	)
	for i := 0; i < 10; i++ {
        wg.Add(1)
        go write(&x,&rwlock,&wg)
    }

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go read(&x,&rwlock,&wg)
    }
    wg.Wait()
    end := time.Now()
    fmt.Println(end.Sub(start),"x ===",x)
}

func main() {
	//互斥锁
	// lockDemo()
	//读写锁
	rwLockDemo()
}