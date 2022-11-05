package main

/*
方法接收者为值和指针的区别:
	1、值方法:	可以通过指针和值调用
	2、指针方法: 只能通过指针来调用，
		但有一个例外，如果某个值是可寻址的(可直接使用 & 操作符取地址的对象，就是可寻址的)，
	    那么编译器会在值调用指针方法时，自动插入取地址符
	
	解释: 指针类型本身就可以调用指针方法，通过*T也可获取对应的值类型，所以同样也可以调用值方法;
		值类型本身则可以调用值方法，如果值类型能通过&T寻址获取到指针，则就可以调用指针方法，如果不能寻址，则不能调用指针方法

	判断是否可寻址: 
		可以被寻址的是左值，既可以出现在赋值号左边也可以出现在右边；
		不可以被寻址的即为右值，比如函数返回值、字面值、常量值等等，只能出现在赋值号右边

方法接收者为值和指针的规则:
	• 类型 T 方法集包含全部 receiver T 方法。
    • 类型 *T 方法集包含全部 receiver T + *T 方法。
    • 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。 
    • 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。 
    • 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法		

值方法和指针方法的取舍: 
	1、方法是否需要修改 receiver 本身。如果需要，那 receiver 必然要是指针了
	2、效率问题。如果 receiver 是值，那在方法调用时一定会产生 struct 拷贝，而大对象拷贝代价很大
	3、一致性。对于同一个 struct 的方法，value method 和 pointer method 混杂用肯定是不优雅的啦
*/

import (
    "fmt"
)

//定义结构体类型
type Foo struct {
    name string
}

//指针方法: 指针类型可以调用，可寻址的值类型也可以调用
func (f *Foo) PointerMethod() {
    fmt.Println("pointer method on", f.name)
}
//值方法: 值类型和指针类型都可以调用
func (f Foo) ValueMethod() {
    fmt.Println("value method on", f.name)
}

//函数创建一个Foo值类型实例，因为函数只能在=右边，也就是右值，右值不能被寻址，并且函数本身也不能被寻址
func NewFooVal() Foo { // 返回一个右值
    return Foo{name: "right value struct"}
}

func NewFooPointer() *Foo { // 返回一个右值
    return &Foo{name: "right value struct"}
}


func main() {
	//创建一个Foo值类型实例，赋给变量f1，变量f1是左值，通过&f1可以被寻址，所以可以调用指针方法
    f1 := Foo{name: "value struct"}
    f1.PointerMethod() // 编译器会自动插入取地址符，变为 (&f1).PointerMethod()
    f1.ValueMethod()
	//创建一个Foo指针类型实例，可以调用值方法和指针方法
    f2 := &Foo{name: "pointer struct"}
    f2.PointerMethod() 
    f2.ValueMethod() // 编译器会自动解引用，变为 (*f2).PointerMethod()

	//函数返回Foo值类型实例，没有左值(变量)，是不能被寻址的，所以只能调用值方法
    NewFooVal().ValueMethod()
    // NewFooVal().PointerMethod() // Error!!!

	//函数返回Foo指针类型实例，本身返回的就是指针类型(指针指向了值)，所以可以调用值方法和指针方法
	NewFooPointer().ValueMethod()
	NewFooPointer().PointerMethod()
}