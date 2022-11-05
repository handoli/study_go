package main

import "fmt"

/*
方法定义:	func (recevier type) methodName(参数列表)(返回值列表){}

一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针,所有给定类型的方法属于该类型的方法集。

Golang 方法总是绑定对象实例（结构体），并隐式将实例作为第一实参 (receiver)
	• 只能为当前包内命名类型（结构体）定义方法。
	• 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
	• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
	• 不支持方法重载，receiver 只是参数签名的组成部分。
	• 可用实例 value 或 pointer 调用全部方法，编译器自动转换。

示例如下:
	// 声明接受者类型:	type Test struct{}
	【 值类型调用方法 】
	// 无参数、无返回值:	func (t Test) method0() {}
	// 单参数、无返回值:	func (t Test) method1(i int) {}
	// 多参数、无返回值:	func (t Test) method2(x, y int) {}
	// 无参数、单返回值:	func (t Test) method3() (i int) { return }
	// 多参数、多返回值:	func (t Test) method4(x, y int) (z int, err error) { return }
	【 指针类型调用方法 】
	// 无参数、无返回值:	func (t *Test) method5() {}
	// 单参数、无返回值:	func (t *Test) method6(i int) {}
	// 多参数、无返回值:	func (t *Test) method7(x, y int) {}
	// 无参数、单返回值:	func (t *Test) method8() (i int) { return }
	// 多参数、多返回值:	func (t *Test) method9(x, y int) (z int, err error) { return }

普通函数与方法的区别:
	1.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然。
	2.对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以。
*/

//定义User结构体类型
type User struct {
    Name  string
    Email string
}

/*
接受者为User[值]类型的方法:	不是同一个对象，是对象的值拷贝
注意:
	函数的接受者是值类型, 即使用指针类型调用, 那么函数内部也是对副本值的操作
*/
func (u User) NotifyVal() {
    fmt.Printf("%v : %v \n", u.Name, u.Email)
	fmt.Printf("Value: %p\n", &u)
}

/*
接受者为User[指针]类型的方法: 是同一个对象
注意:
	函数的接受者是指针类型, 即使用值类型调用, 那么函数内部也是对指针的操作
*/
func (u *User) NotifyRef() {
    fmt.Printf("%v : %v \n", u.Name, u.Email)
	fmt.Printf("Pointer: %p\n", u)
}

func main() {
    
    u1 := User{"月宫", "yuegong@163.com"}
	fmt.Printf("Data: %p\n", u1)
	u2 := &u1
	fmt.Printf("Data: %p\n", u2)

	// [值]类型调用---接受者[值]类型方法
	u1.NotifyVal()
	// [值]类型调用---接受者[指针]类型方法
	u1.NotifyRef()
	// [指针]类型调用---接受者[值]类型方法
	u2.NotifyVal()
	// [指针]类型调用---接受者[指针]类型方法
	u2.NotifyRef()
} 
