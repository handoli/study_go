package main

import (
	"fmt"
	"math"
)

/*
匿名函数是指不需要定义函数名的一种函数实现方式。
在Go里面，函数可以像普通变量一样被传递或使用，Go语言支持随时在代码里定义匿名函数。
匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
*/

func main() {
	// 定义匿名函数，并赋值给变量，变量名就等同于函数名称
	getSqrt := func(a float64) float64 {
        return math.Sqrt(a)
    }
    fmt.Println(getSqrt(4))

	// 匿名函数后面跟'（值）',就相当于调用了函数，此时变量就是接受函数的返回值
	reslt := func(a float64) float64 {
		return math.Sqrt(a)
	}(4)
	fmt.Println(reslt)
		
}