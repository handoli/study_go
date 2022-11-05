package main

import "fmt"

/*
Go语言的For循环有3中形式，只有其中的一种使用分号
	for init; condition; post { }
    for condition { }
    for { }

	init： 一般为赋值表达式，给控制变量赋初值；
	condition： 关系表达式或逻辑表达式，循环控制条件；
	post： 一般为赋值表达式，给控制变量增量或减量

	例一：
		s := "abc"
		for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
			println(s[i])
		}
	例二：
		n := len(s)
		for n > 0 {                // 替代 while (n > 0) {}
			n--
			println(s[n])        // 替代 for (; n > 0;) {}
		}
	例三：
		for {                    // 替代 while (true) {}
			println(s)            // 替代 for (;;) {}
		}

for语句执行过程如下：
	①先对表达式 init 赋初值；
	②判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

*/

func main() {
	//基本使用
    //例一
	var a int
    for a = 0; a < 10; a++ {
       fmt.Printf("a 的值为: %d\n", a)
    }
	fmt.Println("--------------------")
	//例二
	var b int = 15
    for a < b {
       a++
       fmt.Printf("a 的值为: %d\n", a)
    }
	fmt.Println("--------------------")
	//例三
	numbers := [6]int{1, 2, 3, 5}
    for i,x:= range numbers {
       fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
    }


	//嵌套for循环，在for循环中嵌套一个或多个for循环
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i/j); j++ {
		   if(i%j==0) {
			  break // 如果发现因子，则不是素数
		   }
		}
		if(j > (i/j)) {
		   fmt.Printf("%d  是素数\n", i)
		}
	}
	
	//无限循环, while(true){}、for(,,)
	// for { //或者写成 for true { }
	// 	fmt.Printf("这是无限循环。\n");
	// }


	//循环语句range,类似迭代器操作，返回(索引,值)或(键,值)。可对slice、map、数组、字符串等进行迭代循环,  注意：在遍历数组时会复制对象，建议使用应用类型的切片、以及其他引用类型对象

	s := "abc"
    for i := range s {
        println(s[i])
    }
    // ‘_’忽略 index。
    for _, c := range s {
        println(c)
    }
    // 忽略全部返回值，仅迭代。
    for range s {

    }

    m := map[string]int{"a": 1, "b": 2}
    // 返回 (key, value)。
    for k, v := range m {
        println(k, v)
    }
}