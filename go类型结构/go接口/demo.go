package main

import (
	"fmt"
)

/*
在Go语言中接口（interface）是一种类型，一种抽象的类型
接口是一组方法的集合。定义了一个对象的行为规范，只定义不实现，由具体的对象来实现规范的细节

接口规范与类型的关系:
	1、接口只有方法声明，没有实现，没有数据字段
    2、接口可以匿名嵌入其他接口，或嵌入到结构中
    3、对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，
	   既无法修改复制品的状态，也无法获取指针
    4、只有当接口存储的类型和对象都为nil时，接口才等于nil
    5、接口调用不会做receiver的自动转换(结构体方法和接口方法的重要区别)
    6、接口同样支持匿名字段方法
    7、接口也可实现类似OOP中的多态
    8、空接口可以作为任何类型数据的容器
    9、一个类型可实现多个接口，一个接口也可被多个类型实现
	10、接口类型变量能够存储所有实现了该接口的实例
	11、类型的方法集中只要拥有该接口'对应的全部方法'签名。就表示它 "实现" 了该接口,
	    所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值

接口定义:
	type 接口类型名 interface{
        方法名1( 参数列表1 ) 返回值列表1
        方法名2( 参数列表2 ) 返回值列表2
        …
    }
	1、使用type将接口定义为自定义的类型名
	2、接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等
    3、方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包之外的代码访问
    4、参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略

接口的值接收者方法和指针接收者方法:
	若定义的接收者是值，则既可以用接口值调用，也可以用接口指针调用
	若定义的接收者是指针，则只能用接口指针调用，不能用接口值调用

空接口:
	1、指没有定义任何方法的接口。因此任何类型都实现了空接口
	2、空接口类型的变量，可以存储任意类型的变量

	空接口的应用:
		1、使用空接口实现可以接收任意类型的函数参数
			func show(a interface{}) {
				fmt.Printf("type:%T value:%v\n", a, a)
			}
		2、使用空接口实现可以保存任意值的字典
			var studentInfo = make(map[string]interface{})
			studentInfo["name"] = "李白"
			studentInfo["age"] = 18
			studentInfo["married"] = false
			fmt.Println(studentInfo)

类型断言:	空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？
	接口值: 接口是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值
	例子:
		var w io.Writer			//动态类型:	nil				动态值: nil
		w = os.Stdout			//动态类型:	*os.File		动态值: os.File指针
		w = new(bytes.Buffer)	//动态类型:	*bytes.Buffer	动态值: bytes.Buffer
		w = nil 				//动态类型:	nil				动态值: nil
	
	判断空接口中的值: 
		x.(T): x：表示类型为interface{}的变量, T：表示断言x可能是的类型
			   该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，
			   若为true则表示断言成功，为false则表示断言失败		
*/

type Mover interface {
    move()
	say()
}

type Animal interface {
	Eating()
}

type Person interface {
	Named()
	Animal
}

type dog struct {} 

type cat struct {}

type user struct {}

//值类型接收者
func (d dog) move() {
    fmt.Println("value type---狗会动")
} 
//指针类型接收者
func (d *dog) say() {
    fmt.Println("ref type----狗会动")
}

func (d *dog) Eating(){
	fmt.Println("dog ref type----Eating")
}

func (c *cat) Eating(){
	fmt.Println("cat ref type----Eating")
}

func (u *user) Named(){
	fmt.Println("Person---Named")
}
func (u *user) Eating(){
	fmt.Println("Person-Animal-Eating")
}

// 空接口作为函数参数
func show(a interface{}) {
    fmt.Printf("type:%T value:%v\n", a, a)
} 
//封装接口类型断言函数
func justifyType(x interface{}) {
    switch v := x.(type) {
    case string:
        fmt.Printf("x is a string，value is %v\n", v)
    case int:
        fmt.Printf("x is a int is %v\n", v)
    case bool:
        fmt.Printf("x is a bool is %v\n", v)
    default:
        fmt.Println("unsupport type！")
    }
} 

func main() {
	//值类型和指针类型的调用区别
	// vd := dog{}
	// var d1 Mover = dog{} //值类型,报错，接口方法的值类型不能寻址，所以相当于没有实现say()方法
    rd := &dog{}
	var d2 Mover = rd  // 指针类型
	d2.move()
	d2.say()

	//一个类型可以实现多个接口
	var d3 Animal = rd
	d3.Eating()
	//多个类型可以实现同一个接口
	var c1 Animal = &cat{}
	c1.Eating()

	//接口嵌套
	u := &user{}
	//实现的外层接口类型，可以接受实现者类型，包含了自身和嵌入的接口方法
	var p1 Person = u
	p1.Named()
	p1.Eating()
	//实现的内层接口类型，可以接受实现者类型，只包含了自身的接口方法
	var a1 Animal = u
	a1.Eating()

	//空接口
    var x interface{} 	// 定义一个空接口x
    s := "pprof.cn"
    x = s	//接口可以接受字符串类型
    fmt.Printf("type:%T value:%v\n", x, x)
    i := 100
    x = i	//接口可以接受数字类型
    fmt.Printf("type:%T value:%v\n", x, x)
    b := true
    x = b	//接口可以接受bool类型
    fmt.Printf("type:%T value:%v\n", x, x)

	show("yuegong") //空接口作为函数参数
	
    var studentInfo = make(map[string]interface{})	// 空接口作为map值
    studentInfo["name"] = "李白"
    studentInfo["age"] = 18
    studentInfo["married"] = false
    fmt.Println(studentInfo) 


	//接口值类型断言
	var y interface{}
    y = "pprof.cn"
    v, ok := y.(string)
    if ok {
        fmt.Println(v)
    } else {
        fmt.Println("类型断言失败")
    }
	//多个断言，封装函数并使用switch语句来实现
	justifyType(y)

	
}