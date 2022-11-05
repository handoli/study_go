package main

import "fmt"

/*
"_"标识符，用来忽略函数的某个返回值
Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。
没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。
*/

//被命名的返回值参数(可看做与形参类似的局部变量),最后由return隐式返回
func add(a, b int) (c int) {
    c = a + b
    return
}

//返回多个参数
func calc(a, b int) (int, int) {
    sum := a + b
    avg := (a + b) / 2
    return sum, avg
}

//示例，把其他函数返回的多个参数当作函数形参
func sum(n ...int) int {
	sum := 0
	for _,v := range n {
		sum += v
	}
	return sum
}

func main() {
    var a, b int = 1, 2
    c := add(a, b)
	//  "_",忽略函数的某个返回值
    _, avg := calc(a, b)
    fmt.Println(a, b, c, avg)

	//函数返回值作为其他函数的形参
	count := sum(calc(a,b))
	fmt.Println(count)
}
