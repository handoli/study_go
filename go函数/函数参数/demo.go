package main

import "fmt"

//值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数
func swapVal(x,y int) {
	temp := x
	x = y
	y = temp
}

//引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数
func swapRef(x,y *int) {
	temp := *x
	*x = *y
	*y = temp
}

//不定参数传值,就是函数的参数不是固定的，后面的类型是固定的
func vals(args ...int){
	for _,v := range args{
		fmt.Println(v)
	}
}
//不定参数传值(无类型),就是函数的参数不是固定的，后面的类型是固定的
func refs(args ...interface{}){
	for _,v := range args{
		fmt.Println(v)
	}
}


func main() {
	var x,y = 1,2
	swapVal(x,y)
	fmt.Println("val: x=",x,"y=",y)
	swapRef(&x,&y)
	fmt.Println("val: x=",x,"y=",y)

	vs := []int{75,43,23,5,7,0,12,9}
	vals(vs...) //使用切片时，必须展开（...）

	refs(1,"h",65,93,"dl")
}