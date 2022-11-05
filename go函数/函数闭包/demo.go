package main

import "fmt"

//函数b嵌套在函数a内部,函数a返回函数b
func a() func() int {
	i := 0
	b := func () int  {
		i++
        fmt.Println(i)
        return i
	}
	return b
}

//返回多个闭包
func test01(base int) (func(int) int, func(int) int) {
    // 定义2个函数，并返回
    // 相加
    add := func(i int) int {
        base += i
        return base
    }
    // 相减
    sub := func(i int) int {
        base -= i
        return base
    }
    // 返回
    return add, sub
}

func main() {
	/*
	变量c实际上是指向了函数b()，再执行函数c()后就会显示i的值，第一次为1，第二次为2，第三次为3，以此类推.
	这段代码就创建了一个闭包。因为函数a()外的变量c引用了函数a()内的函数b(),
	由于闭包的存在使得函数a()返回后，a中的i始终存在，这样每次执行c()，i都是自加1后的值
	*/
	c := a()
	c()
	c()
	c()
	/*
	c()跟c2()引用的是不同的环境，在调用i++时修改的不是同一个i，因此两次的输出都是1。
	函数a()每进入一次，就形成了一个新的环境，对应的闭包中，函数都是同一个函数，环境却是引用不同的环境。这和c()和c()的调用顺序都是无关的
	*/
	c2 := a()
	c2()
	c2()

	//返回多个闭包函数
	f1,f2 := test01(10)
	//base一直是没有消,并且被当前环境的f1、f2共享 
	fmt.Println(f1(1), f2(2))
    fmt.Println(f1(3), f2(4))
}