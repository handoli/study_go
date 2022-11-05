package main

import "fmt"

/*
嵌套匿名成员: 一个结构体内嵌套了另一个匿名的结构体（java中对象复合的概念）
	被嵌套的匿名结构体类型的方法，可以被外层结构体调用，由编译器负责查找
	同时，通过这种嵌套组合，实现了复用和“override”效果
*/

//内层结构体
type User struct {
    id   int
    name string
}

type Manager struct {
    User	//匿名的嵌套结构体类型
    title string
}

//内层结构体的方法，实现“override”效果
func (self *User) ToString() string {
    return fmt.Sprintf("User: %p, %v", self, self)
}

//外层结构体方法
func (self *Manager) ToString() string {
    return fmt.Sprintf("Manager: %p, %v", self, self)
}

func main() {
	//声明，变量m指向外层结构体
    m := Manager{User{1, "Tom"}, "Administrator"}
	//调用外层结构体的方法
	//注意：如果外层结构体类型没有此方法，则由编译器去查找内部匿名结构体的此方法
    fmt.Println(m.ToString())
	//调用内部匿名结构体的方法
    fmt.Println(m.User.ToString())
}  