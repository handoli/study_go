package main

import "fmt"

/*
方法表达式：说简单点，其实就是方法对象赋值给变量,根据调用者不同，方法分为两种表现形式
    1、method value: instance.method(args...)
		通过struct实例(值类型实例、指针类型实例)获取方法对象,struct实例就是方法的接受者,隐式调用(绑定实例),
		方法接受者为值类型时会复制receiver(与结构体的值方法和指针方法概念一样)
	2、method expression: <type>.func(instance, args...)
		通过struct类型(类似于java中的类,struct类型也分为值类型和指针类型)获取方法对象,在调用时需要传递struct实例对象作为参数(显示调用)
		重点: struct值类型和指针类型也符合方法集规则
			1、struct值类型只能获取 值类型接受者方法,在调用时要传递值类型实例
			2、struct指针类型可以获取值类型和指针类型的接受者方法，在调用时 要传递指针类型实例

*/

type Student struct {
    id   int
    name string
}

func (s *Student) SkillPointer() {
    fmt.Printf("指针型函数:%p, %v\n", s, s)
}

func (s Student) SkillValue() {
    fmt.Printf("值类型函数: %p, %v\n", &s, s)
}

func main() {
    s := Student{1, "乔帮主"} // 结构体实例化
	sp := &s
    //常规使用方式
    s.SkillPointer()
	s.SkillValue()
    fmt.Println(".............................\n")

    /*
	方法值
	struct实例就是方法的接受者,隐式调用(绑定实例)
	*/
	//值类型实例调用 指针类型接受者方法，因为是可寻址的，所以可以调用
    vp := s.SkillPointer //这个就是方法值，调用函数时，无需再传递接收者，隐藏了接收者
    vp()
	//值类型实例调用 值类型接受者方法，值类型接受者方法会复制receiver
    vv := s.SkillValue
    vv()
	//指针类型实例调用 指针类型接受者方法
	pp := sp.SkillPointer
	pp()
	//指针类型实例调用 值类型接受者方法，值类型接受者方法会复制receiver
	pv := sp.SkillValue
	pv()
    fmt.Println(".............................\n")

	/*
	方法表达式
	通过struct类型获取方法对象,在调用时需要传递struct实例对象作为参数(显示调用)
	*/
	//struct指针类型可以获取值类型和指针类型的接受者方法，在调用时 要传递指针类型实例
    tpp := (*Student).SkillPointer  // struct指针类型获取指针类型接受者方法
    tpp(&s)	//调用时要传递指针类型实例
	tpv := (*Student).SkillValue    // struct指针类型获取值类型接受者方法
    tpv(&s)	//调用时要传递指针类型实例
	//struct值类型只能获取 值类型接受者方法,在调用时要传递值类型实例
    tvv := Student.SkillValue 
    tvv(s) //在调用时要传递值类型实例
}