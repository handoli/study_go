package main

import "fmt"

//参数就是函数的定义
func test(fn func() int) int {
    return fn()
}

//下面与上面的区别是，上面是匿名定义函数，下面是先单独定义函数
// 函数类型定义
type FormatFunc func(s string, x, y int) string 
//接受指定类型定义的函数
func format(fn FormatFunc, s string, x, y int) string {
    return fn(s, x, y)
}

func main() {
	//调用时给予函数的实现
    s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。

	//调用时给予函数的实现
    s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

    println(s1, s2)
}