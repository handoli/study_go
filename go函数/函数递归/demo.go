package main

import "fmt"

/*
递归，就是在运行的过程中调用自己。
一个函数调用自己，就叫做递归函数。

构成递归需具备的条件：
 	1.子问题须与原始问题为同样的事，且更为简单。
    2.不能无限制地调用本身，须有个出口，化简为非递归状况处理。
*/

//数字阶乘
func factorial(i int) int {
    if i <= 1 {
        return 1
    }
    return i * factorial(i-1)
}

//斐波那契数列(Fibonacci)
func fibonaci(i int) int {
    if i == 0 {
        return 0
    }
    if i == 1 {
        return 1
    }
    return fibonaci(i-1) + fibonaci(i-2)
}


func main() {
	//数字阶乘
	fmt.Println(factorial(7))

	//斐波那契数列
	for i := 0; i<10; i++ {
		fmt.Println(fibonaci(i))
	}
}