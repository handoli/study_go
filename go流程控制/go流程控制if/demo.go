package main

import "fmt"

func main() {
	
	x := 10

	if x > 0 {
		fmt.Println("if条件判断表达式")
	}

	//if...else if...else
	if str := "abcde"; x > 0 {
		fmt.Println(str[1])
	}else if x < 0 {
		fmt.Println(str[2])
	}else {
		fmt.Println(str[0])
	}

	//嵌套if
	y := 20	
	if x > 0 {
		if y > x {
			fmt.Println("y - x = ",y - x)
		} else {
			fmt.Println("x - y = ",x - y)
		}
	}else {
		fmt.Println("x + y = ",x + y)
	}

	//逻辑判断表达式
	if x > 0 && y > 0 {
		fmt.Println("我们都大于0")
	}else if x < 0 || y < 0 {
		fmt.Println("我们其中有小于0")
	}else {
		fmt.Println("可能不存在")
	}
}