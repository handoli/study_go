package main

import "fmt"

/*
select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。具体用法查看go并发编程Select
*/
func main() {
	//基本用法
	var c1, c2, c3 chan int
    var i1, i2 int
    select {
        case i1 = <-c1:
        	fmt.Printf("received ", i1, " from c1\n")
        case c2 <- i2:
        	fmt.Printf("sent ", i2, " to c2\n")
        case i3, ok := <-c3:  // same as: i3, ok := <-c3
        	if ok {
            	fmt.Printf("received ", i3, " from c3\n")
        	} else {
            	fmt.Printf("c3 is closed\n")
        	}
        default:
        	fmt.Printf("no communication\n")
    }
	
	//典型用法,判断超时、退出、判断channel是否阻塞

}